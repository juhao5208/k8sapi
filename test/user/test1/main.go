package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
 * @author  巨昊
 * @date  2021/11/6 14:56
 * @version 1.15.3
 */

type user struct {
	id       string
	uuid     string
	username string
	password string
	usertype string
	usertel  string
	useradd  string
}

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql",
		"root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
	if err != nil {
		panic(err)
	}
	db = database
}

func getLen(db *sqlx.DB) int {
	len := 0
	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		len++
	}
	return len
}

func getAll(db *sqlx.DB, length int) [3]user {
	data := [3]user{}
	rows, _ := db.Query("select * from user")
	i := 0
	for rows.Next() {
		err := rows.Scan(&data[i].id, &data[i].uuid, &data[i].username, &data[i].password,
			&data[i].usertype, &data[i].usertel, &data[i].useradd)
		if err != nil {
			panic(err)
		}
		i++
	}
	return data
}

func main() {
	//rows, _ := db.Query("select * from user")
	//len := getLen(db)
	//line := data{}
	//for rows.Next() {
	//	err := rows.Scan(&line.id, &line.uuid, &line.username, &line.password,
	//		&line.usertype, &line.usertel, &line.useradd)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//fmt.Println(line)

	len := getLen(db)
	data := getAll(db, len)
	fmt.Println(data)
}
