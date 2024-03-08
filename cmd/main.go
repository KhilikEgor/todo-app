package main

import (
	"github.com/KhilikEgor/todo-app"
	"github.com/KhilikEgor/todo-app/pkg/hendler"
	"github.com/KhilikEgor/todo-app/pkg/repository"
	"github.com/KhilikEgor/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := hendler.NewHendler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running http server: %s", err.Error())
	}
}
