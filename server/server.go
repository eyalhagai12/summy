package server

import (
	"summy/workerpool"

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

	return server
}

func (s *Server) Start() error {
	return s.app.Start(":8080")
}
