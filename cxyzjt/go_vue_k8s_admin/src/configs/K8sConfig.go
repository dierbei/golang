package configs

import (
	"log"

	"go_vue_k8s_admin/src/services"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sConfig struct {
	DeployHandler    *services.DeployHandler    `inject:"-"`
	PodHandler       *services.PodHandler       `inject:"-"`
	NameHandler      *services.NameSpaceHandler `inject:"-"`
	EventHandler     *services.EventHandler     `inject:"-"`
	IngressHandler   *services.IngressHandler   `inject:"-"`
	ServiceHandler   *services.ServiceHandler   `inject:"-"`
	SecretHandler    *services.SecretHandler    `inject:"-"`
	ConfigMapHandler *services.ConfigMapHandler `inject:"-"`
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
	factory.Core().V1().Namespaces().Informer().AddEventHandler(cfg.NameHandler)
	factory.Core().V1().Events().Informer().AddEventHandler(cfg.EventHandler)
	factory.Networking().V1beta1().Ingresses().Informer().AddEventHandler(cfg.IngressHandler)
	factory.Core().V1().Services().Informer().AddEventHandler(cfg.ServiceHandler)
	factory.Core().V1().Secrets().Informer().AddEventHandler(cfg.SecretHandler)
	factory.Core().V1().ConfigMaps().Informer().AddEventHandler(cfg.ConfigMapHandler)

	factory.Start(wait.NeverStop)
	return factory
}
