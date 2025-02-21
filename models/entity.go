package models

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
