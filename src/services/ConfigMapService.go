package services

import (
	"k8s.io/client-go/kubernetes"
	"k8sapi2/src/models"
)

/**
 * @author  巨昊
 * @date  2021/9/8 17:00
 * @version 1.15.3
 */

//@service
type ConfigMapService struct {
	Client    *kubernetes.Clientset `inject:"-"`
	ConfigMap *ConfigMapStruct      `inject:"-"`
}

func NewConfigMapService() *ConfigMapService {
	return &ConfigMapService{}
}

//前台用于显示Secret列表
func (this *ConfigMapService) ListConfigMap(ns string) []*models.ConfigMapModel {
	list := this.ConfigMap.ListAll(ns)
	ret := make([]*models.ConfigMapModel, len(list))
	for i, item := range list {
		ret[i] = &models.ConfigMapModel{
			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
		}
	}
	return ret
}
