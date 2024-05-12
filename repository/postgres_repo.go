package repository

import (
	"context"
	"database/sql"

	db "github.com/Prthmesh6/crud_app/db/sqlc"
	"github.com/Prthmesh6/crud_app/models"
)

type taskRepo struct {
	db *db.Queries
}

func NewTaskRepo(db *db.Queries) TaskRepository {
	return &taskRepo{
		db: db,
	}
}

func (t *taskRepo) Add(ctx context.Context, task models.Task) (taskId int, err error) {
	taskParams := db.CreateTaskParams{
		Name:        task.Name,
		Description: sql.NullString{String: task.Description},
		Status:      int32(task.Status),
		Priority:    int32(task.Priority),
		DueDate:     int64(task.DueDate),
	}
	tsk, err := t.db.CreateTask(ctx, taskParams)
	taskId = int(tsk.ID)
	return
}

func (t *taskRepo) Update(ctx context.Context, task models.Task) (err error) {
	params := db.UpdateTaskStatusParams{ID: task.ID, Status: int32(task.Status)}
	err = t.db.UpdateTaskStatus(ctx, params)
	return
}

func (t *taskRepo) GetTasksList(ctx context.Context, limit, offset int) (tasks []models.Task, err error) {
	params := db.ListTasksParams{Limit: int32(limit), Offset: int32(offset)}
	dbTasks, err := t.db.ListTasks(ctx, params)

	for _, val := range dbTasks {
		t := models.Task{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description.String,
			Status:      int(val.Status),
			Priority:    int(val.Priority),
			DueDate:     int(val.DueDate),
		}
		tasks = append(tasks, t)
	}
	return
}
