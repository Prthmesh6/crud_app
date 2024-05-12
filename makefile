postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=prathmesh -d postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todo_app

dropdb:
	docker exec -it postgres12 dropdb todo_app

migrateup:
	migrate -path="db/migration" -database="postgresql://root:prathmesh@localhost:5432/todo_app?sslmode=disable" -verbose up

migratedown:
	migrate -path="db/migration" -database="postgresql://root:prathmesh@localhost:5432/todo_app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc