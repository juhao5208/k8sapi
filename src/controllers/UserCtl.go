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
	var data = make([]models.User, length)
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

func (this *UserCtl) Build(goft *goft.Goft) {
	goft.Handle("POST", "/vue-admin-template/user/login", this.login)
	goft.Handle("POST", "/vue-admin-template/user/logout", this.logout)
	goft.Handle("GET", "/vue-admin-template/user/info", this.info)
	goft.Handle("GET", "/user/list", this.getList)
	goft.Handle("DELETE", "/user/list", this.deleteUser)
}
