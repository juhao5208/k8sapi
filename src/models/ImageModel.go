package models

/**
 * @author  巨昊
 * @date  2021/11/8 15:36
 * @version 1.15.3
 */

//定义镜像结构体
type Image struct {
	//repository string `json:"repository"`
	//tag        string `json:"tag"`
	ID         string            `json:"id"`
	Containers int64             `json:"containers"`
	Created    string            `json:"created"`
	Size       int64             `json:"size"`
	Labels     map[string]string `json:"labels"`
}
