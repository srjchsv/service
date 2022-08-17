package services

import (
	"github.com/srjchsv/service/internal/repository"
)

type Authorization interface {
	CreateUser(user repository.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	RefreshToken(token string, userID int) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization)}
}
