package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Prthmesh6/crud_app/models"
	"github.com/Prthmesh6/crud_app/service"
)

type Middleware struct {
	taskService service.TaskService
}

func New(service service.TaskService) *Middleware {
	return &Middleware{
		taskService: service,
	}
}

func (th *Middleware) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID, err := th.taskService.Add(r.Context(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"task_id": taskID})
}

func (th *Middleware) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = th.taskService.Update(r.Context(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *Middleware) GetTasksListHandler(w http.ResponseWriter, r *http.Request) {
	orderBy, err := strconv.Atoi(r.URL.Query().Get("order_by"))
	if err != nil {
		http.Error(w, "Invalid order_by parameter", http.StatusBadRequest)
		return
	}

	// Validating orderBy parameter
	if orderBy != 0 && orderBy != 1 {
		http.Error(w, "Invalid order_by parameter", http.StatusBadRequest)
		return
	}

	tasks, err := th.taskService.GetTasksList(r.Context(), orderBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (th *Middleware) MarkTaskDoneHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := r.URL.Query().Get("task_id")
	if taskIDStr == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	taskID, err := strconv.ParseInt(taskIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	err = th.taskService.MarkTaskDone(r.Context(), int(taskID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
