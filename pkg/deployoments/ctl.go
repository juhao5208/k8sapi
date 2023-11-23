package deployoments

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi2/src/services"
)

/**
 * @author  巨昊
 * @date  2021/8/21 16:35
 * @version 1.15.3
 */

type DeploymentCtlV2 struct {
	K8sClient *kubernetes.Clientset   `inject:"-"`
	DeployMap *services.DeploymentMap `inject:"-"`
}

func NewDeploymentCtlV2() *DeploymentCtlV2 {
	return &DeploymentCtlV2{}
}

//快捷创建时  需要 初始化一些 标签
func (this *DeploymentCtlV2) initLabel(deploy *v1.Deployment) {
	if deploy.Spec.Selector == nil {
		deploy.Spec.Selector = &metav1.LabelSelector{MatchLabels: map[string]string{"jtapp": deploy.Name}}
	}
	if deploy.Spec.Selector.MatchLabels == nil {
		deploy.Spec.Selector.MatchLabels = map[string]string{"jtapp": deploy.Name}
	}
	if deploy.Spec.Template.ObjectMeta.Labels == nil {
		deploy.Spec.Template.ObjectMeta.Labels = map[string]string{"jtapp": deploy.Name}
	}
	deploy.Spec.Selector.MatchLabels["jtapp"] = deploy.Name
	deploy.Spec.Template.ObjectMeta.Labels["jtapp"] = deploy.Name
}

func (this *DeploymentCtlV2) SaveDeployment(c *gin.Context) goft.Json {
	dep := &v1.Deployment{}
	goft.Error(c.ShouldBindJSON(dep))
	if c.Query("fast") != "" { //代表是快捷创建 。 要预定义一些值
		this.initLabel(dep)
	}
	update := c.Query("update") //代表是更新
	if update != "" {
		_, err := this.K8sClient.AppsV1().Deployments(dep.Namespace).Update(c, dep, metav1.UpdateOptions{})
		goft.Error(err)
	} else {
		_, err := this.K8sClient.AppsV1().Deployments(dep.Namespace).Create(c, dep, metav1.CreateOptions{})
		goft.Error(err)
	}

	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (this *DeploymentCtlV2) RmDeployment(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")

	err := this.K8sClient.AppsV1().Deployments(ns).Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (this *DeploymentCtlV2) LoadDeploy(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	dep, err := this.DeployMap.GetDeployment(ns, name) // 原生
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": dep,
	}
}

func (this *DeploymentCtlV2) Build(goft *goft.Goft) {
	//路由
	goft.Handle("GET", "/deployments/:ns/:name", this.LoadDeploy)
	goft.Handle("POST", "/deployments", this.SaveDeployment)
	//删除deploy
	goft.Handle("DELETE", "/deployments/:ns/:name", this.RmDeployment)
}
func (*DeploymentCtlV2) Name() string {
	return "DeploymentCtlV2"
}
