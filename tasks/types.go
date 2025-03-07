package tasks

type GetAllTasksRequest struct{}

type InProgressTasksPageRequest struct {
	Page int `query:"page"`
	Size int `query:"size"`
}
