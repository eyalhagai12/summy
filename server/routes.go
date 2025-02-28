package server

import (
	"net/http"
	"summy/api"
	"summy/source/email"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes(apiGroup *echo.Group) {
	gmailAuthHandlers := email.NewGmailAuthHandlers(s.db, s.oauthConfig)
	apiGroup.GET("/gmail/auth/callback", api.HandlerFromFunc(gmailAuthHandlers.GetAuthCode, http.StatusOK))
	apiGroup.POST("/gmail/auth", api.HandlerFromFunc(gmailAuthHandlers.HandleGmailAuth, http.StatusOK))
}
