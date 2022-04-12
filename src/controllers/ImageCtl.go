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

/*func ListAllImage() {
	//获取client
	defaultHeaders := map[string]string{
		"User-Agent": "engine-api-cli-1.0",
	}
	cli, err := client.NewClient("tcp://192.168.99.100:2375",
		"v1.41", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}
	//获取本机所有的容器
	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}
	//打印容器id
	for _, c := range containers {
		fmt.Println(c.ID)
	}
}*/

type ImageCtl struct {
	ImageService *services.ImageService `inject:"-"` //首字母一定要大写
}

func (*ImageCtl) Name() string {
	return "ImageCtl"
}

func NewImageCtl() *ImageCtl {
	return &ImageCtl{}
}

/*//配置ssh相关
type Cli struct {
	IP         string      //IP地址
	Username   string      //用户名
	Password   string      //密码
	Port       int         //端口号
	client     *ssh.Client //ssh客户端
	LastResult string      //最近一次Run的结果
}

//创建命令行对象
//@param ip IP地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New(ip string, username string, password string, port ...int) *Cli {
	cli := new(Cli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}
	return cli
}

//执行shell
//@param shell shell脚本命令
func (c Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

//连接
func (c *Cli) connect() error {
	config := ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	c.client = sshClient
	return nil
}*/

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
	/*cli := New("192.168.99.100", "docker", "tcuser", 22)
	output, err := cli.Run("docker images")
	goft.Error(err)
	return gin.H{
		"code": "20000",
		"data": output,
	}*/
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
