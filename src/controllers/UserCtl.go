package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi2/src/models"
	"k8sapi2/src/services"
	"log"
)

type UserCtl struct {
	UserService *services.UserService `inject:"-"` //首字母一定要大写
}

func (*UserCtl) Name() string {
	return "UserCtl"
}

func NewUserCtl() *UserCtl {
	return &UserCtl{}
}

func (this *UserCtl) login(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": gin.H{
			"token": "admin-token",
		},
	}
}

func (this *UserCtl) logout(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (this *UserCtl) info(c *gin.Context) string {
	c.Header("Content-type", "application/json")
	return `{"code":20000,"data":{"roles":["admin"],
		"introduction":"I am a super administrator","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif","name":"Super Admin"}}`

}

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

func (this *UserCtl) getList(c *gin.Context) goft.Json {
	c.Header("Content-type", "application/json")
	len := getLen(db)
	list := getAll(db, len)

	return gin.H{
		"code": 20000,
		"data": list,
	}
}

func (this *UserCtl) deleteUser(c *gin.Context) goft.Json {
	c.Header("Content-type", "application/json")
	name := c.DefaultQuery("name", "")
	result, err := db.Exec("delete from user where username=?", name)
	if err != nil {
		fmt.Println("ERROR", err)
		log.Fatal(err)
	}
	x, _ := result.LastInsertId()
	fmt.Println(x)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
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

func getAll(db *sqlx.DB, length int) []models.User {
	var data = make([]models.User,length)
	rows, _ := db.Query("select * from user")
	i := 0
	for rows.Next() {
		err := rows.Scan(&data[i].UserId, &data[i].UUID, &data[i].Username, &data[i].Password,
			&data[i].UserType, &data[i].UserTel, &data[i].UserAdd)
		if err != nil {
			panic(err)
		}
		i++
	}
	return data
}

//func GetQueryColumns(rows *sql.Rows) ([]string, map[string]string, error) {
//	columnTypes, err := rows.ColumnTypes()
//	if err != nil {
//		return nil, nil, err
//	}
//	length := len(columnTypes)
//	columns := make([]string, length)
//	columnTypeMap := make(map[string]string, length)
//	for i, ct := range columnTypes {
//		columns[i] = ct.Name()
//		columnTypeMap[ct.Name()] = ct.DatabaseTypeName()
//	}
//	return columns, columnTypeMap, nil
//}
//
//func QueryForInterface(db *sql.DB, sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
//	rows, err := db.Query(sqlInfo, args...)
//	defer rows.Close()
//	if err != nil {
//		return nil, err
//	}
//	columns, columnTypeMap, err := GetQueryColumns(rows)
//	if err != nil {
//		return nil, err
//	}
//	columnLength := len(columns)
//	cache := make([]interface{}, columnLength)
//	for index, _ := range cache {
//		var a interface{}
//		cache[index] = &a
//	}
//	var list []map[string]interface{} //返回的切片
//	for rows.Next() {
//		_ = rows.Scan(cache...)
//		item := make(map[string]interface{})
//		for i, data := range cache {
//			if ct, ok := columnTypeMap[columns[i]]; ok {
//				if (ct == "VARCHAR" || ct == "DATETIME") && *data.(*interface{}) != nil {
//					item[columns[i]] = string((*data.(*interface{})).([]byte))
//				} else {
//					item[columns[i]] = *data.(*interface{})
//				}
//			} else {
//				item[columns[i]] = *data.(*interface{})
//			}
//		}
//		list = append(list, item)
//	}
//	return list, nil
//}

/*//获取所有用户
func (this *UserCtl) GetList(c *gin.Context) goft.Json {
	sqlString := "select * from user"
	userJson, err := getJSON(sqlString)
	if err != nil {
		log.Fatal(err)
	}
	return userJson
}

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
}*/

//func (this *UserCtl) GetList(c *gin.Context) goft.Json {
//	var user []models.Users
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

////数据库指针
//var db *sqlx.DB
//
////初始化数据库连接，init()方法系统会在动在main方法之前执行。
//func init() {
//	database, err := sqlx.Open("mysql", "root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//	}
//	db = database
//}

func (this *UserCtl) Build(goft *goft.Goft) {
	goft.Handle("POST", "/vue-admin-template/user/login", this.login)
	goft.Handle("POST", "/vue-admin-template/user/logout", this.logout)
	goft.Handle("GET", "/vue-admin-template/user/info", this.info)
	goft.Handle("GET", "/user/list", this.getList)
	goft.Handle("DELETE", "/user/list", this.deleteUser)
}
