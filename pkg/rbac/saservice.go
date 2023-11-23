package rbac

import corev1 "k8s.io/api/core/v1"

/**
 * @author  巨昊
 * @date  2021/9/17 16:20
 * @version 1.15.3
 */

//@Service
type SaService struct {
	SaMap *SaMapStruct `inject:"-"`
}

func NewSaService() *SaService {
	return &SaService{}
}

func (this *SaService) ListSa(ns string) []*corev1.ServiceAccount {
	return this.SaMap.ListAll(ns)
}
