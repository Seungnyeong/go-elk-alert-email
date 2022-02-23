package service

import (
	"test/keyinfo/db"
	"test/keyinfo/domain"
)

type UserRepository struct {
	db db.DataBase
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db.NewMysqlDatabase()}
}

func (repo UserRepository) FindAdminUser() ([]domain.User, error) {
	result, err := repo.db.FindAdminUser()
	return result, err
}

func (repo UserRepository) FindUser(username string) (domain.User, error) {
	result, err := repo.db.FindUser(username)
	return result, err
}
