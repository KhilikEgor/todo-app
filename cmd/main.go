package main

import (
	"github.com/KhilikEgor/todo-app"
	"github.com/KhilikEgor/todo-app/pkg/hendler"
	"github.com/KhilikEgor/todo-app/pkg/repository"
	"github.com/KhilikEgor/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing cingings: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := hendler.NewHendler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
