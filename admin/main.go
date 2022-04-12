package admin
//
//import (
//	"context"
//	"flag"
//	"fmt"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/tools/clientcmd"
//	"os"
//	"time"
//)
//
///**
// * @author  巨昊
// * @date  2021/8/12 19:29
// * @version 1.15.3
// */
//
////func main(){
////	config ,err:=clientcmd.BuildConfigFromFlags("","C:\\Users\\巨昊\\.kube\\config")
////	if err!=nil{
////		panic(err)
////	}
////	fmt.Println(config.Host)
////}
//
//func main() {
//	// 配置 k8s 集群外 kubeconfig 配置文件
//	var kubeconfig *string
//	kubeconfig = flag.String("config", "C:\\Users\\巨昊\\.kube\\config", "absolute path to the kubeconfig file")
//	flag.Parse()
//
//	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// 根据指定的 config 创建一个新的 clientset
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		panic(err.Error())
//	}
//	for {
//		// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string)返回 PodInterface
//		// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
//		// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 Pod 列表
//		pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
//		if err != nil {
//			panic(err.Error())
//		}
//		fmt.Printf("There are %d pods in the k8s cluster\n", len(pods.Items))
//
//		// 获取指定 namespace 中的 Pod 列表信息
//		namespace := "default"
//		pods, err = clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("\nThere are %d pods in namespaces %s\n", len(pods.Items), namespace)
//		for _, pod := range pods.Items {
//			fmt.Printf("Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
//		}
//		time.Sleep(10 * time.Second)
//	}
//}
//
//func prettyPrint(maps map[string]interface{}) {
//	lens := 0
//	for k, _ := range maps {
//		if lens <= len(k) {
//			lens = len(k)
//		}
//	}
//	for key, values := range maps {
//		spaces := lens - len(key)
//		v := ""
//		for i := 0; i < spaces; i++ {
//			v += " "
//		}
//		fmt.Printf("%s: %s%v\n", key, v, values)
//	}
//}
//
//func homeDir() string {
//	if h := os.Getenv("HOME"); h != "" {
//		return h
//	}
//	return os.Getenv("USERPROFILE") // windows
//}
