package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	db "github.com/Prthmesh6/crud_app/db/sqlc"
	"github.com/Prthmesh6/crud_app/models"
	"github.com/Prthmesh6/crud_app/repository"
	"github.com/Prthmesh6/crud_app/service"
	_ "github.com/lib/pq"
)

func main() {

	//This is hardcoded, this will be taken from env variables or CONSUL (Hashicorp)
	conn, err := sql.Open("postgres", "postgresql://root:prathmesh@localhost:5432/todo_app?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	Queries := db.New(conn)
	taskRespo := repository.NewTaskRepo(Queries)
	todoService := service.New(taskRespo)

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

}
