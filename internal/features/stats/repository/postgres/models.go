package statistics_postgres_repository

import (
	"time"

	"github.com/Aam-Shaegar/todo-list/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}

func taskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	taskDomains := make([]domain.Task, len(taskModels))
	for i, task := range taskModels {
		taskDomains[i] = taskDomainFromModel(task)
	}
	return taskDomains
}
