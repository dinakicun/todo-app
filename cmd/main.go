package main

import (
	"log"

	todo "github.com/dinakicun/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
