package main

import (
"zoomer/db"
"database/sql"
)

func main() {
database, err := sql.Open("postgres", "postgres://user:password@localhost/mydb?sslmode=disable")
if err != nil {
panic(err)
}
defer database.Close()

redisClient := db.NewRedisClient()

entityRepo := repository.NewRedisRepository(database, redisClient)
entityUsecase := usecase.NewEntityUsecase(entityRepo, redisClient)
}
