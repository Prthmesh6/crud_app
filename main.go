package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "github.com/Prthmesh6/crud_app/db/sqlc"
	"github.com/Prthmesh6/crud_app/models"
	"github.com/Prthmesh6/crud_app/repository"
	"github.com/Prthmesh6/crud_app/router"
	"github.com/Prthmesh6/crud_app/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:8080"
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

	a, err := todoService.Add(context.Background(), models.Task{
		Name:        "Task1",
		Description: "random",
		Status:      1,
		Priority:    1,
		DueDate:     int(time.Now().Unix()),
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
	tList, _ := todoService.GetTasksList(context.TODO(), 1)
	fmt.Printf("taskList is %+v", tList)

	taskHandler := router.New(todoService)

	r := mux.NewRouter()

	// Define routes and their respective handlers
	r.HandleFunc("/tasks/add", taskHandler.AddTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/update", taskHandler.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/list", taskHandler.GetTasksListHandler).Methods("GET")
	r.HandleFunc("/tasks/mark_done", taskHandler.MarkTaskDoneHandler).Methods("PUT")

	// Start server
	fmt.Println("Starting server at 8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//-----Graceful Shutdown------

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	go func() {
		log.Printf("Server started on %v \n", defaultRoutingServiceURL)
		log.Println(server.ListenAndServe())
	}()

	<-s
	shutDown(server)

}

func shutDown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Handle error while server shutdown")
	}
	log.Println("doing gracefull shutdown ")
}
