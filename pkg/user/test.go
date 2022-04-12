package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//
//import (
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//	"k8sapi2/src/models"
//)
//
///**
// * @author  巨昊
// * @date  2021/10/28 20:44
// * @version 1.15.3
// */
//
////数据库指针
//var db *sqlx.DB
//
////用户结构体
//type Users struct {
//	UserId string `db:"id"`
//	//CreateTime time.Time `db:"created_at"`
//	//UpdateTime time.Time `db:"updated_at"`
//	//DeleteTime time.Time `db:"deleted_at"`
//	UUID     string `db:"uuid"`
//	Username string `db:"username"`
//	Password string `db:"password"`
//	UserType string `db:"usertype"`
//	UserTel  string `db:"usertel"`
//	UserAdd  string `db:"useradd"`
//}
//
////初始化数据库连接，init()方法系统会在动在main方法之前执行。
//func init() {
//	database, err := sqlx.Open("mysql", "root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//	}
//	db = database
//}
//
//func ListAll() {
//	rows, err := db.Query("select * from user")
//	if err != nil {
//		panic(err)
//	}
//	for rows.Next() {
//		var id, uuid, username, password, usertype, usertel, useradd string
//		err = rows.Scan(&id, &uuid, &username, &password, &usertype, &usertel, &useradd)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(id, uuid, username, password, usertype, usertel, useradd)
//	}
//}
//
//func GetList() [...]models.Users {
//	var user [...]models.Users
//	rows, err := db.Query("select * from user")
//	if err != nil {
//		panic(err)
//	}
//	for rows.Next() {
//		//var id, uuid, username, password, usertype, usertel, useradd string
//		var i int = 0
//		err = rows.Scan(&user[i].UserId, &user[i].UUID, &user[i].Username, &user[i].Password, &user[i].UserType,
//			&user[i].UserTel, &user[i].UserAdd)
//		if err != nil {
//			panic(err)
//		}
//		//fmt.Println(id, uuid, username, password, usertype, usertel, useradd)
//		i++
//	}
//	/*return gin.H{
//		"code": 20000,
//		"data": &models.Users{
//			UserId:   user[0].UserId,
//			UUID:     user[0].UUID,
//			Username: user[0].Username,
//			Password: user[0].Password,
//			UserType: user[0].UserType,
//			UserTel:  user[0].UserTel,
//			UserAdd:  user[0].UserAdd,
//		},
//	}*/
//
//	return user
//
//}
//
//func getUserById(user_id string) {
//	var users []Users
//	sql := "select id, username,usertel,usertype,useradd from user where id=? "
//	err := db.Select(&users, sql, user_id)
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		return
//	}
//	fmt.Println("select succ:", users)
//}
//
//func UpdateUserById(user_id string) {
//	//执行SQL语句
//	sql := "update user set username =? where id = ?"
//	res, err := db.Exec(sql, "李四", 2)
//
//	if err != nil {
//		fmt.Println("exec failed,", err)
//		return
//	}
//
//	//查询影响的行数，判断修改插入成功
//	row, err := res.RowsAffected()
//	if err != nil {
//		fmt.Println("rows failed", err)
//	}
//
//	fmt.Println("update succ:", row)
//}
//
//func DeleteUserById(user_id string) {
//
//	sql := "delete from user where id=?"
//
//	res, err := db.Exec(sql, 2)
//	if err != nil {
//		fmt.Println("exce failed,", err)
//		return
//	}
//
//	row, err := res.RowsAffected()
//	if err != nil {
//		fmt.Println("row failed, ", err)
//	}
//	fmt.Println("delete succ: ", row)
//}
//
//func main() {
//	/*sql := "insert into user(username,password,usertype, usertel,useradd)values (?,?,?,?,?)"
//	value := [...]string{"lisi", "admin", "administrator", "12643597561", "北京市"}
//
//	//执行SQL语句
//	r, err := db.Exec(sql, value[0], value[1], value[2], value[3], value[4])
//	if err != nil {
//		fmt.Println("exec failed,", err)
//		return
//	}
//
//	//查询最后一天用户ID，判断是否插入成功
//	id, err := r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed,", err)
//		return
//	}
//	fmt.Println("insert succ", id)
//
//	getUserById("123456")*/
//
//	//ListAll()
//
//	u := GetList()
//	fmt.Println(u[0].Username)
//
//}

//将mysql数据中查找的数据转换为json格式
func getJSON(sqlString string) (string, error) {
	//sqlString为sql语句
	var db *sqlx.DB
	database, err := sqlx.Open("mysql", "root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
	if err != nil {
		panic(err)
	}
	db = database

	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}

func main() {
	data, err := getJSON("select * from user")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
