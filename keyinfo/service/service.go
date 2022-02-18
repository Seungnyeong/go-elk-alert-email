package service

import (
	"fmt"
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

func (repo UserRepository) FindAdminUser() []domain.User  {
	result, err := repo.db.FindAdminUser()
	utils.CheckError(err)
	fmt.Println(result)
	return result
}

func (repo UserRepository) FindUser(username string) domain.User {
	result, err := repo.db.FindUser(username)
	utils.CheckError(err)
	return result
}