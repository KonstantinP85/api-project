package repository

import (
	"github.com/KonstantinP85/api-project"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user api_project.User) (int, error)
}

type ApiList interface {

}

type Repository struct {
	Authorization
	ApiList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
