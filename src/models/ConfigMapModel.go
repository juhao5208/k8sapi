package models

/**
 * @author  巨昊
 * @date  2021/9/8 16:59
 * @version 1.15.3
 */

//提交 用的
type PostConfigMapModel struct {
	Name      string
	NameSpace string
	Data      map[string]string
	IsUpdate  bool
}

//列表用
type ConfigMapModel struct {
	Name       string
	NameSpace  string
	CreateTime string
	Data       map[string]string // KV

}
