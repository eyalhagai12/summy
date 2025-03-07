package tasks

import (
	"fmt"
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
	return templates.Home()
}

func (tv *TaskViews) InProgressTasks(c echo.Context, request InProgressTasksPageRequest) templ.Component {
	offset := (request.Page - 1) * request.Size

	fmt.Printf("offset: %d", offset)

	tasks := []models.Task{}
	err := tv.db.Select(&tasks, "SELECT * FROM tasks LIMIT $2 OFFSET $1", offset, request.Size)
	if err != nil {
		return templates.InProgressTasks([]models.Task{}, 0)
	}

	return templates.InProgressTasks(tasks, request.Page)
}
