package server

import (
	"log"
	"summy/workerpool"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type Server struct {
	app         *echo.Echo
	db          *sqlx.DB
	oauthConfig *oauth2.Config
	wp          *workerpool.WorkerPool
}

func New(config Config) *Server {
	db := sqlx.MustConnect("postgres", config.DbConnectionString)
	oauthConfig := oauth2.Config{
		ClientID:     config.GmailClientID,
		ClientSecret: config.GmailClientSecret,
		RedirectURL:  config.HostURL + "/api/gmail/auth/callback",
		Scopes:       []string{gmail.GmailReadonlyScope},
		Endpoint:     google.Endpoint,
	}
	app := echo.New()

	server := &Server{
		app:         app,
		db:          db,
		oauthConfig: &oauthConfig,
		wp:          workerpool.New(config.WorkerPoolSize, config.TaskBufferPerWorker),
	}

	server.RegisterRoutes(app.Group("/api"))
	server.RegisterViews(app)

	_, err := db.Exec(
		"INSERT INTO tasks (id, title, description, status, user_id, source) VALUES ($1, $2, $3, $4, $5, $6);",
		uuid.New(),
		"test task",
		"this is the test taks i use",
		"discovered",
		uuid.MustParse("679fb22e-a314-4561-85d4-574d34eca9b1"),
		"manual",
	)
	if err != nil {
		log.Println("failed to add random user -  " + err.Error())
	}

	return server
}

func (s *Server) Start() error {
	return s.app.Start(":8080")
}
