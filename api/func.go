package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpHandlerFunc[Req any, Res any] func(echo.Context, Req) (Res, error)

func HandlerFromFunc[Req any, Res any](handler HttpHandlerFunc[Req, Res], successStatus int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request Req
		if err := bindToRequest(c, &request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request - ", err)
		}

		response, err := handler(c, request)
		if err != nil {
			return err
		}

		return c.JSON(successStatus, response)
	}
}
