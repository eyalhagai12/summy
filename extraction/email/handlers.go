package email

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type GmailAuthHandlers struct {
	db          *sqlx.DB
	oauthConfig *oauth2.Config
}

func NewGmailAuthHandlers(db *sqlx.DB, oauthconfig *oauth2.Config) GmailAuthHandlers {
	return GmailAuthHandlers{
		db:          db,
		oauthConfig: oauthconfig,
	}
}

func (g *GmailAuthHandlers) GetAuthCode(c echo.Context, request GmailAuthCodeRequest) (string, error) {
	ctx := c.Request().Context()

	tok, err := g.oauthConfig.Exchange(ctx, request.Code)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "failed to exchange oauth token")
	}

	authInfo := GmailAuthInformation{
		UserID:       uuid.New(),
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
	}

	_, err = sqlx.NamedExec(
		g.db,
		"INSERT INTO gmail_auth_info (user_id, access_code, refresh_token, expiration) VALUES ($1, $2, $3, $4);",
		authInfo,
	)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "failed to store auth information - ", err)
	}

	return "", nil
}
