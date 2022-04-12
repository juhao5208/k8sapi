package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"k8sapi2/src/models"
)

/**
 * @author  巨昊
 * @date  2021/11/12 6:45
 * @version 1.15.3
 */

//初始化docker客户端
func NewClient() (*client.Client, error) {
	return client.NewClient("tcp://192.168.99.100:2376", "v1.41", nil, nil)
}

//获取镜像数量
func getImageNum() int {
	num := 0 //镜像数量
	cli, err := NewClient()
	if err != nil {
		fmt.Println("创建docker客户端失败", err)
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}
	num = len(images)
	return num
}

func main() {
	/*num := getImageNum()
	fmt.Println(num)*/

	len := getImageNum()
	var data = make([]models.Image, len)
	cli, _ := NewClient()
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}
	for i := 0; i < len; i++ {
		data[i].ID = images[i].ID
		data[i].Containers = images[i].Containers
		//data[i].Created = images[i].Created
		data[i].Size = images[i].Size
		data[i].Labels = images[i].Labels
	}
	fmt.Println(data)
}
