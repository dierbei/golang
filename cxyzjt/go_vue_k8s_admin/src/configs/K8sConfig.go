package configs

import (
	"log"

	"go_vue_k8s_admin/src/core"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sConfig struct {
	DeployHandler *core.DeployHandler `inject:"-"`
	PodHandler    *core.PodHandler    `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

// InitClient 初始化K8s客户端
func (cfg *K8sConfig) InitClient() *kubernetes.Clientset {
	config := &rest.Config{
		Host: "http://175.24.198.168:8009",
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// InitInformer 初始化Informer监听
func (cfg *K8sConfig) InitInformer() informers.SharedInformerFactory {
	factory := informers.NewSharedInformerFactory(cfg.InitClient(), 0)

	factory.Apps().V1().Deployments().Informer().AddEventHandler(cfg.DeployHandler)
	factory.Core().V1().Pods().Informer().AddEventHandler(cfg.PodHandler)

	factory.Start(wait.NeverStop)
	return factory
}
