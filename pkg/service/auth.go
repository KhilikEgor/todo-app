package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/KhilikEgor/todo-app"
	"github.com/KhilikEgor/todo-app/pkg/repository"
)

const salt = "hjfwaoifj4685fza6dnbvjkdnjkv"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
