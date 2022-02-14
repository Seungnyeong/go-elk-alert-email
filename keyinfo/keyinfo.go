package keyinfo

import (
	"database/sql"
	"fmt"
	"test/utils"

	_ "github.com/go-sql-driver/mysql"
)

const (
	database = "keyinfo_nd"
	user = "secutech"
	password = "qhdksxla123!@#"
	host = "127.0.0.1"
)



type KeyInfo struct {
	Username	string `json:"username"`
	Email		string `json:"email"`
	Key			sql.NullString `json:"key"`
}

func GetKeysWithUser() {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	db, err := sql.Open("mysql", connectionString)	
	utils.CheckError(err)
	defer db.Close()

	err = db.Ping()
	utils.CheckError(err)
	fmt.Println("connect sucess")

	
	

	rows, err := db.Query("select username, email, k.key from auth_user left join keyinfo k on auth_user.id = k.user_id")
	utils.CheckError(err)
	var u KeyInfo
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&u.Username, &u.Email, &u.Key)
		utils.CheckError(err)
		if len(u.Key.String) > 0 {
			fmt.Printf("Data row = (%s, %s, %s)\n", u.Username, u.Email, u.Key.String)
		}
	}
	err = rows.Err()
	utils.CheckError(err)
}