package controllers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi2/src/models"
	"k8sapi2/src/services"
	"time"
)

/**
 * @author  巨昊
 * @date  2021/11/8 15:38
 * @version 1.15.3
 */

type ImageCtl struct {
	ImageService *services.ImageService `inject:"-"` //首字母一定要大写
}

func (*ImageCtl) Name() string {
	return "ImageCtl"
}

func NewImageCtl() *ImageCtl {
	return &ImageCtl{}
}

//初始化docker客户端
func NewClient() (*client.Client, error) {
	return client.NewClient("tcp://172.17.199.167:8443", "v1.25.2", nil, nil)
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

func getImage(len int) []models.Image {
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

		var timestamp int64 = images[i].Created
		timeobj := time.Unix(int64(timestamp), 0)
		date := timeobj.Format("2006-01-02 15:04:05")
		data[i].Created = date
		data[i].Size = images[i].Size
		data[i].Labels = images[i].Labels
	}
	return data
}

func (this *ImageCtl) getList(c *gin.Context) goft.Json {
	c.Header("Content-type", "application/json")
	num := getImageNum()
	list := getImage(num)
	return gin.H{
		"code": 20000,
		"data": list,
	}
}

func (this *ImageCtl) deleteImage(c *gin.Context) goft.Json {
	c.Header("Content-type", "application/json")
	id := c.DefaultQuery("id", "")
	cli, err := NewClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	removeImage, err := cli.ImageRemove(ctx, id, types.ImageRemoveOptions{})
	if err != nil {
		return err
	}
	fmt.Println(removeImage)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (this *ImageCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/image/list", this.getList)
	goft.Handle("DELETE", "/image/list", this.deleteImage)
}
