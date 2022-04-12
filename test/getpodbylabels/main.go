package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

/**
 * @author  巨昊
 * @date  2021/10/25 15:07
 * @version 1.15.3
 */

var clientset *kubernetes.Clientset

const ns = "rdp"

////构建容器模型
//type Container struct{
//	name string
//}

func main() {
	var username string
	fmt.Print("输入用户名：")
	fmt.Scanf("%s", &username)
	fmt.Println(username)

	//加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\巨昊\\.kube\\config")
	if err != nil {
		panic(err)
	}
	//创建客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//获取pod列表
	podlist, err := clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{
		//LabelSelector: fmt.Sprintf("app=%s", username),
	})
	if err != nil {
		panic(err)
	}

	for _, pod := range podlist.Items {
		if username == pod.Labels["user"] {
			fmt.Println(pod.Name, pod.Spec.Containers)
		}
	}
}