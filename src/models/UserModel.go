package models

/**
 * @author  巨昊
 * @date  2021/11/1 17:01
 * @version 1.15.3
 */

//用户结构体
type User struct {
	UserId     string    `json:"id"`
	UUID       string    `json:"uuid"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	UserType   string    `json:"usertype"`
	UserTel    string    `json:"usertel"`
	UserAdd    string    `json:"useradd"`
}
