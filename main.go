package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "github.com/Prthmesh6/crud_app/db/sqlc"
	"github.com/Prthmesh6/crud_app/repository"
	"github.com/Prthmesh6/crud_app/router"
	"github.com/Prthmesh6/crud_app/service"
	_ "github.com/lib/pq"
)

func main() {

	//This is hardcoded, this will be taken from env variables or CONSUL (Hashicorp)
	dbdriver, dbUrl := "postgres", "postgresql://root:prathmesh@localhost:5432/todo_app?sslmode=disable"
	conn, err := sql.Open(dbdriver, dbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	//created a connection as passed as dependency,
	//down the line we can also change this database
	//this is the advantage of making system decoupled.
	Queries := db.New(conn)
	taskRespo := repository.NewTaskRepo(Queries)

	//Here I created todo service by passing one db
	//we can create another usermanagement service
	//with different database without distubring current implementation
	todoService := service.New(taskRespo)

	//Now using this service I can call the business logic.
	//As I passed DB as dependency, similarly one can pass Logger & instrumenting object

	taskHandler := router.New(todoService)

	// Define routes and their respective handlers
	http.HandleFunc("/tasks/add", taskHandler.AddTaskHandler)
	http.HandleFunc("/tasks/update", taskHandler.UpdateTaskHandler)
	http.HandleFunc("/tasks/list", taskHandler.GetTasksListHandler)
	http.HandleFunc("/tasks/mark_done", taskHandler.MarkTaskDoneHandler)

	// Start server
	fmt.Println("Starting server at 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
