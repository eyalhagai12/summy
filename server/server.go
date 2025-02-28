package server

import (
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
}

func New(config Config) *Server {
	db := sqlx.MustConnect("postgres", config.DbConnectionString)
	oauthConfig := oauth2.Config{
		ClientID:     config.GmailClientID,
		ClientSecret: config.GmailClientSecret,
		RedirectURL:  config.ServerURL + "/api/gmail/auth/callback",
		Scopes:       []string{gmail.GmailReadonlyScope},
		Endpoint:     google.Endpoint,
	}
	app := echo.New()

	server := &Server{
		app:         app,
		db:          db,
		oauthConfig: &oauthConfig,
	}

	server.RegisterRoutes(app.Group("/api"))

	return server
}

func (s *Server) Start() error {
	return s.app.Start(":8080")
}
