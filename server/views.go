package server

import (
	"net/http"
	"summy/api"
	"summy/tasks"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterViews(app *echo.Echo) {
	taskViews := tasks.NewTaskViews(s.db, s.wp)
	app.GET("/", api.ViewFromFunc(taskViews.TasksHome, http.StatusOK))
}
