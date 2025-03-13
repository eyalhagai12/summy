package tasks

import (
	"fmt"
	"net/url"
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

func (tv *TaskViews) TasksHome(c echo.Context, _ any) templ.Component {
	return templates.Home()
}

func (tv *TaskViews) TaskList(c echo.Context, request PaginatedTasksRequest) templ.Component {
	offset := (request.Page - 1) * request.Size
	nextPageURL := c.Request().URL
	queryParams := url.Values{}
	queryParams.Add("size", fmt.Sprintf("%d", request.Size))
	queryParams.Add("status", request.Status)
	nextPageURL.RawQuery = queryParams.Encode()

	tasks := []models.Task{}
	err := tv.db.Select(&tasks, "SELECT * FROM tasks WHERE status = $3 LIMIT $2 OFFSET $1", offset, request.Size, request.Status)
	if err != nil {
		fmt.Println("error: %w", err)
		return templates.PaginatedTaskList([]models.Task{}, 0, 1, nextPageURL)
	}

	return templates.PaginatedTaskList(tasks, request.Page, request.Size, nextPageURL)
}

func (tv *TaskViews) AddTaskModal(c echo.Context, _ any) templ.Component {
	return templates.AddTaskModal()
}
