package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"test/config"
	"test/keyinfo/domain"
	"test/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DataBase interface {
	FindAdminUser() ([]domain.User, error)
	FindUser(username string) (domain.User, error)
}

type MySqlConnection struct {
	url string
}

func NewMySqlConnection(url string) *MySqlConnection {
	return &MySqlConnection{url}
}

type MysqlDatabase struct {
	client *sql.DB
}

var errCannotFindUser = errors.New("there is NOT that user")

func NewMysqlDatabase() *MysqlDatabase {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		config.P.Database.User,
		config.P.Database.Password,
		config.P.Database.Host,
		config.P.Database.Name,
	)

	client, err := sql.Open("mysql", connectionString)
	if err != nil {
		utils.CheckError(err)
	}

	client.SetConnMaxLifetime(time.Second * 10)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return &MysqlDatabase{client}
}

func (ms MysqlDatabase) FindAdminUser() ([]domain.User, error) {
	users := make([]domain.User, 0)
	rows, err := ms.client.Query("select username, email, is_superuser, is_active from auth_user where is_superuser = true AND is_active = true")
	if err != nil {
		return nil, err
	}
	utils.CheckError(err)
	defer ms.client.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Username, &user.Email, &user.IsSuperUser, &user.IsActive)
		if err != nil {
			log.Fatal("Error while scanning keyinfo " + err.Error())
			return nil, err
		}

		users = append(users, user)

	}
	return users, nil
}

func (ms MysqlDatabase) FindUser(username string) (domain.User, error) {
	var user domain.User
	err := ms.client.QueryRow("select username, email, is_superuser, is_active from auth_user where username = ?", username).Scan(&user.Username, &user.Email, &user.IsSuperUser, &user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
			return user, errCannotFindUser
		} else {
			log.Fatal(err)
		}
	}
	return user, err
}
