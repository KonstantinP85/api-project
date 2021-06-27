package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/KonstantinP85/api-project"
	"github.com/KonstantinP85/api-project/pkg/repository"
)

const salt = "neuicnoennjsdcnsc"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user api_project.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}