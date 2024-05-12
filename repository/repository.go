package repository

import (
	"context"

	"github.com/Prthmesh6/crud_app/models"
)

type UserRepository interface {
	Add(user models.User)
	Delete(id int)
	Update(id int, name string, email string)
	GetAllUsers() []models.User
}

type TaskRepository interface {
	Add(ctx context.Context, task models.Task) (taskId int, err error)
	Update(ctx context.Context, task models.Task) (err error)
	GetTasksList(ctx context.Context, limit, offset int) (tasks []models.Task, err error)
}
