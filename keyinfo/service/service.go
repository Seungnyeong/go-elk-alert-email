package service

import (
	"test/keyinfo/db"
	"test/keyinfo/domain"
	"test/utils"
)

type UserRepository struct {
	db db.DataBase
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db.NewMysqlDatabase()}
}

func (repo UserRepository) FindAdminUser() ([]domain.User, error ) {
	result, err := repo.db.FindAdminUser()
	utils.CheckError(err)
	return result, err
}

func (repo UserRepository) FindUser(username string) (domain.User, error) {
	result, err := repo.db.FindUser(username)
	return result, err
}