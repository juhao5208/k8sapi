package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi2/src/services"
)

/**
 * @author  巨昊
 * @date  2021/7/18 14:39
 * @version 1.15.3
 */

type NsCtl struct {
	NsMap *services.NsMapStruct `inject:"-"`
}

func NewNsCtl() *NsCtl {
	return &NsCtl{}
}

func (this *NsCtl) ListAll(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": this.NsMap.ListAll(),
	}
}
func (this *NsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/nslist", this.ListAll)
}
func (*NsCtl) Name() string {
	return "NsCtl"
}
