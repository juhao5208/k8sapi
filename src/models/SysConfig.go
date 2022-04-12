package models

/**
 * @author  巨昊
 * @date  2021/9/14 9:46
 * @version 1.15.3
 */

type ClusterInfo struct {
	EndPoint string `yaml:"endpoint"`
	CaFile   string `yaml:"cafile"`
	UserCert string `yaml:"usercert"` //用户证书存放位置
}
type NodesConfig struct {
	Name string
	Ip   string
	User string
	Pass string
}
type K8sConfig struct {
	Nodes       []*NodesConfig
	ClusterInfo *ClusterInfo `yaml:"cluster-info"`
}
type SysConfig struct {
	K8s *K8sConfig
}
