package tasks

import (
	"summy/models"
	"summy/templates"
	"summy/workerpool"

	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type TaskViews struct {
	db *sqlx.DB
	wp *workerpool.WorkerPool
}

func NewTaskViews(db *sqlx.DB, wp *workerpool.WorkerPool) *TaskViews {
	return &TaskViews{
		db: db,
		wp: wp,
	}
}

func (tv *TaskViews) TasksHome(c echo.Context, request GetAllTasksRequest) templ.Component {
	tasks := []models.Task{}
	err := tv.db.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		return templates.Home(tasks, err)
	}

	return templates.Home(tasks, nil)
}
