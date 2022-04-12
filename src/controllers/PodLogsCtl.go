package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"time"
)

/**
 * @author  巨昊
 * @date  2021/9/9 21:46
 * @version 1.15.3
 */

type PodLogsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewPodLogsCtl() *PodLogsCtl {
	return &PodLogsCtl{}
}
func (this *PodLogsCtl) GetLogs(c *gin.Context) (v goft.Void) {
	ns := c.DefaultQuery("ns", "default")
	podname := c.DefaultQuery("podname", "")
	cname := c.DefaultQuery("cname", "")
	var tailLine int64 = 100
	opt := &v1.PodLogOptions{Follow: true, Container: cname, TailLines: &tailLine}

	cc, _ := context.WithTimeout(c, time.Minute*30) //设置半小时超时时间。否则会造成内存泄露
	req := this.Client.CoreV1().Pods(ns).GetLogs(podname, opt)
	reader, err := req.Stream(cc)
	goft.Error(err)
	defer reader.Close()
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)       // 如果 当前日志 读完了。 会阻塞
		if err != nil && err != io.EOF { //一旦超时 会进入 这个程序 ,此时一定要break 掉
			break
		}
		w, err := c.Writer.Write([]byte(string(buf[0:n])))
		if w == 0 || err != nil {
			break
		}
		c.Writer.(http.Flusher).Flush()
	}

	return

}
func (*PodLogsCtl) Name() string {
	return "PodLogsCtl"
}

func (this *PodLogsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods/logs", this.GetLogs)
}
