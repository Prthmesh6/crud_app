package models

type TaskStatus int

const (
	Pending TaskStatus = iota
	Done
	Cancelled
)

type Status struct {
	ID     int32 `json:"id"`
	Status int   `json:"status"`
}

type Task struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Priority    int    `json:"priority"`
	DueDate     int    `json:"due_date"`
}

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
