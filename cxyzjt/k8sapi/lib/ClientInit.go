package lib

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)
var K8sClient *kubernetes.Clientset
func init() {
	config:=&rest.Config{
		Host:"http://175.24.198.168:9090",
		BearerToken:"kubeconfig-user-r8wp5.c-sld9n:hdkwqrwrblpmwkgj7m29vjgw7xqmsb2hqwnm7b9rgfgg7xdtft8lqr",
	}
	c,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}
	K8sClient=c
}