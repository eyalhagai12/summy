package tasks

type PaginatedTasksRequest struct {
	Page   int    `query:"page"`
	Size   int    `query:"size"`
	Status string `query:"status"`
}
