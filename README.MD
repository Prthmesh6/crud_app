## SetUp:-
1. Clone the repo (Golang 1.21+, sqlc should be installed)
2. run make postgres
3. run make createdb
4. run make migrateup
5. run make sqlc


## API Information :-

### API :- "http://localhost:8080/tasks/add"
POST 
{
  "name": "Complete Project 3",
  "description": "Finish the coding part of the project",
  "status": 0,
  "priority": 2,
  "due_date": 1630368000
}

### Points :- 
1. When you are including priority in payload, it will work
2. Try it removing priority, it will not work, currently it is saying invalid args, we can add specific message to it later.

### API :- "http://localhost:8080/tasks/lists?order_by=1"
GET

### Points :- 
1. No payload is required for this
2. This returns the list of all tasks
3. Once we add userservice, we can handle multiple users and return tasks for specific user
4. Also we can add more API parameters to get specific tasks like, completed, cancelled
5. Here I am passing order_by param in query only that too as int, this can be done from client. 

### API :- "http://localhost:8080/tasks/update"
PUT 
{
  "id":2
  "name": "Complete Project 3",
  "status": 0
}

### Note :-
This add,update words in URL should not be used, But I have intentionally kept it for now so that it will be easy to understand what api is doing for evaluating
Currently we have given facillity to edit status only of a task. 

### API :- "http://localhost:8080/tasks/mark_done?task_id=1"
PUT 
This will mark task as done


### Additional Points :- 
1. Database versioning also have been added here, so we can keep track of all the changes in DB also we can migrate up and down as per our requirement
2. If any new column needs to be added in db, just add it in DDL, and then run make sqlc, all the code will get generated automatically which helps to maintain a code.
3. I have developed this as per dependency injection stadard, so that you can pass other databse also easily from main.go or when initialising the service. Which makes all services decoupled. 
4. You can also visiit below mentioned repos developed by me to get more information.

- [Microservice](https://github.com/Prthmesh6/Microservice_Gokit_Sample)
- [LRU Cache](https://github.com/Prthmesh6/LRU_Cache)

## Documentation

- [SQLC](https://docs.sqlc.dev/en/latest/)
- [mux router](https://github.com/gorilla/mux)

