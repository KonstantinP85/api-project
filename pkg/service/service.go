package service

import (
	"github.com/KonstantinP85/api-project"
	"github.com/KonstantinP85/api-project/pkg/repository"
)

type Authorization interface {
	CreateUser(user api_project.User) (int, error)
}

type ApiList interface {

}

type Service struct {
	Authorization
	ApiList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}