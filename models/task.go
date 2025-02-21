package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Entity

	Name        string    `db:"name"`
	Description string    `db:"description"`
	DueDate     time.Time `db:"due_date"`
	Status      string    `db:"status"`
	UserID      uuid.UUID `db:"user_id"`
	Source      string    `db:"source"`
}
