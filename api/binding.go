package api

import (
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

func bindToRequest[Req any](c echo.Context, request *Req) error {
	noBodyMethods := []string{http.MethodGet, http.MethodDelete}
	if slices.Contains(noBodyMethods, c.Request().Method) {
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to load request - "+err.Error())
		}

		if err := (&echo.DefaultBinder{}).BindPathParams(c, request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to load request - "+err.Error())
		}

		return nil
	} else {
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to load request - "+err.Error())
		}

		return nil
	}
}
