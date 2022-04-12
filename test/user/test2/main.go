package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
 * @author  巨昊
 * @date  2021/11/8 16:04
 * @version 1.15.3
 */

var db *sql.DB

func deleteUser(name string, db *sql.DB) {
	result, err := db.Exec("delete from user where username=?", name)
	if err != nil {
		fmt.Println("ERROR", err)
		log.Fatal(err)
	}
	x, _ := result.LastInsertId()
	fmt.Println(x)
}

func main() {
	database, err := sql.Open("mysql",
		"root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
	if err != nil {
		panic(err)
	}
	db = database
	deleteUser("lisi", db)
}
