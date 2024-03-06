package main

import (
	"github.com/KhilikEgor/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occurred while running http server: %s", err.Error())
	}
}
