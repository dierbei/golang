package configs

import (
	"go_vue_k8s_admin/pkg/rbac"
	"io/ioutil"
	"log"

	"go_vue_k8s_admin/src/models"
	"go_vue_k8s_admin/src/services"

	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type K8sConfig struct {
	DeployHandler             *services.DeployHandler         `inject:"-"`
	PodHandler                *services.PodHandler            `inject:"-"`
	NameHandler               *services.NameSpaceHandler      `inject:"-"`
	EventHandler              *services.EventHandler          `inject:"-"`
	IngressHandler            *services.IngressHandler        `inject:"-"`
	ServiceHandler            *services.ServiceHandler        `inject:"-"`
	SecretHandler             *services.SecretHandler         `inject:"-"`
	ConfigMapHandler          *services.ConfigMapHandler      `inject:"-"`
	NodeHandler               *services.NodeHandler           `inject:"-"`
	RoleHandler               *rbac.RoleHandler               `inject:"-"`
	RoleBindingHandler        *rbac.RoleBindingHandler        `inject:"-"`
	ServiceAccountHandler     *rbac.ServiceAccountHandler     `inject:"-"`
	ClusterRoleHandler        *rbac.ClusterRoleHandler        `inject:"-"`
	ClusterRoleBindingHandler *rbac.ClusterRoleBindingHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (*K8sConfig) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	config.Insecure = true
	if err != nil {
		log.Fatal(err)
	}
	return config
}

// InitMetricClient metric客户端
func (cfg *K8sConfig) InitMetricClient() *versioned.Clientset {
	c, err := versioned.NewForConfig(cfg.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// InitClient 初始化K8s客户端
func (cfg *K8sConfig) InitClient() *kubernetes.Clientset {
	c, err := kubernetes.NewForConfig(cfg.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// InitSysConfig 初始化节点配置
func (*K8sConfig) InitSysConfig() *models.SysConfig {
	b, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := &models.SysConfig{}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
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
	factory.Core().V1().Nodes().Informer().AddEventHandler(cfg.NodeHandler)
	factory.Rbac().V1().Roles().Informer().AddEventHandler(cfg.RoleHandler)
	factory.Rbac().V1().RoleBindings().Informer().AddEventHandler(cfg.RoleBindingHandler)
	factory.Core().V1().ServiceAccounts().Informer().AddEventHandler(cfg.ServiceAccountHandler)
	factory.Rbac().V1().ClusterRoles().Informer().AddEventHandler(cfg.ClusterRoleHandler)
	factory.Rbac().V1().ClusterRoleBindings().Informer().AddEventHandler(cfg.ClusterRoleBindingHandler)

	factory.Start(wait.NeverStop)
	return factory
}
