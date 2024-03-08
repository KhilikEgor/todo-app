package main

import (
	"github.com/KhilikEgor/todo-app"
	"github.com/KhilikEgor/todo-app/pkg/hendler"
	"log"
)

func main() {
	handlers := new(hendler.Hendler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running http server: %s", err.Error())
	}
}
