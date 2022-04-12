package configs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

/**
 * @author  巨昊
 * @date  2021/11/1 20:42
 * @version 1.15.3
 */



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
