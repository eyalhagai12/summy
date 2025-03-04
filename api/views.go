package api

import (
	"net/http"
	"summy/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type viewFunc[Req any] func(echo.Context, Req) templ.Component

func renderHtml(c echo.Context, component templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	component = templates.Layout("Summy", component)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func ViewFromFunc[Req any](vf viewFunc[Req], successCode int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request Req
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to load request - "+err.Error())
		}

		component := vf(c, request)

		return renderHtml(c, component)
	}
}
