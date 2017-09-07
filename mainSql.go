package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)

const (
	DB_HOST = ""
	DB_NAME = ""
	DB_USER = ""
	DB_PASS = ""
)

type User struct {
	Id string
	Name string
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME))
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("select uid as id, uname as name from user where deleted_yn=?", "N")
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()

	user := new(User)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(user)
	}
}
