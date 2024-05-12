package service

import (
	"context"
	"errors"
	"sort"

	"github.com/Prthmesh6/crud_app/models"
	"github.com/Prthmesh6/crud_app/repository"
)

var (
	ErrInvalidArgument = errors.New("InvalidArguments")
)

type TaskService interface {
	Add(ctx context.Context, task models.Task) (taskId int, err error)
	Update(ctx context.Context, task models.Task) (err error)
	GetTasksList(ctx context.Context, orderBy int) (tasks []models.Task, err error)
}

type taskService struct {
	db repository.TaskRepository
}

func New(db repository.TaskRepository) TaskService {
	return &taskService{
		db: db,
	}
}

func (t *taskService) Add(ctx context.Context, task models.Task) (taskId int, err error) {
	if task.Name == "" || task.Priority < 0 {
		err = ErrInvalidArgument
		return
	}
	taskId, err = t.db.Add(ctx, task)
	return
}

func (t *taskService) Update(ctx context.Context, task models.Task) (err error) {
	if task.ID <= 0 || task.Name == "" || task.Priority < 0 {
		err = ErrInvalidArgument
		return
	}

	err = t.db.Update(ctx, task)
	return
}

func (t *taskService) GetTasksList(ctx context.Context, orderBy int) (tasks []models.Task, err error) {
	//currently not using hardcoding this limit & offset
	// we can take this from api
	tasks, err = t.db.GetTasksList(ctx, 100, 0)

	//Also this logic we can keep in db layer, we will do order by as per sorting request
	if orderBy == 1 {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Priority > tasks[j].Priority
		})
	} else {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Priority > tasks[j].Priority
		})
	}
	return
}

func (t *taskService) MarkTaskDone(ctx context.Context, taskId int) (err error) {
	if taskId < 0 {
		return ErrInvalidArgument
	}
	err = t.db.Update(ctx, models.Task{ID: int64(taskId), Status: int(models.Done)})
	return
}

//we can add methods like
// GetCancelledTasks
//MarkTaskCancelled
//GetCompletedTasks
