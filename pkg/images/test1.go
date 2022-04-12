package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

/**
 * @author: 巨昊
 * @date 2021/11/1 17:18
 * @version 1.17
 */

//func connectMysql() *sqlx.DB {
//	Db, err := sqlx.Open("mysql", "root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
//	if err != nil {
//		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
//	}
//	return Db
//}
//
//func queryData(Db *sqlx.DB) {
//	rows, err := Db.Query("select * from user")
//	if err != nil {
//		fmt.Printf("query faied, error:[%v]", err.Error())
//		return
//	}
//	for rows.Next() {
//		//定义变量接收查询数据
//		var id, uuid, username, password, usertype, usertel, useradd string
//		var createTime, updateTime, deleteTime time.Time
//
//		err := rows.Scan(&id, &createTime,&updateTime, &deleteTime, &uuid, &username, &password, &usertype, &usertel, &useradd)
//		if err != nil {
//			fmt.Println("get data failed, error:[%v]", err.Error())
//		}
//		fmt.Println(username, id, usertel, useradd, usertype)
//	}
//	//关闭结果集（释放连接）
//	rows.Close()
//}
//
//func main() {
//	var Db *sqlx.DB = connectMysql()
//	queryData(Db)
//}

//数据库指针
var db *sqlx.DB

//初始化数据库连接，init()方法系统会在动在main方法之前执行。
func init() {
	database, err := sqlx.Open("mysql", "root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

func main() {
	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, uuid, username, password, usertype, usertel, useradd string
		var createTime, updateTime, deleteTime time.Time
		err = rows.Scan(&id, &createTime, &updateTime, &deleteTime, &uuid, &username,
			&password, &usertype, &usertel, &useradd)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, createTime, updateTime, deleteTime, uuid, username, password, usertype, usertel, useradd)
	}
}
