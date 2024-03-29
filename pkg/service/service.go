package service

import (
	"github.com/KhilikEgor/todo-app"
	"github.com/KhilikEgor/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int error, err error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
