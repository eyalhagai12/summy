package email

import "github.com/google/uuid"

type GmailAuthCodeRequest struct {
	State string `query:"state-token"`
	Code  string `query:"code"`
	Scope string `query:"scope"`
}

type GmailAuthRequest struct {
	UserID string `json:"userId"`
}

type GmailAuthResponse struct {
	URL string `json:"url"`
}

type GmailAuthInformation struct {
	UserID       uuid.UUID `db:"user_id"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
}
