-- +goose Up
-- +goose StatementBegin
CREATE TABLE gmail_auth_info (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) NOT NULL,
    access_token VARCHAR(100) NOT NULL,
    refresh_token VARCHAR(100) NOT NULL,
    expiration TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE gmail_auth_info;
-- +goose StatementEnd
