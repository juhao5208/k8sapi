package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

/**
 * @author  巨昊
 * @date  2021/11/4 16:33
 * @version 1.15.3
 */

func NewPool() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "5208juhao"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "k8sapi2"
	dsn := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil { //ping测试网络连通性以及用户密码是否正确
		log.Fatal(err)
	}
	return db
}
