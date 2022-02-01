package configs

import (
	"go_vue_k8s_admin/pkg/rbac"
	"go_vue_k8s_admin/src/services"

	"k8s.io/client-go/kubernetes"
)

type K8sMaps struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}

//InitDeploymentMap 初始化Deployment存储Map
func (m *K8sMaps) InitDeploymentMap() *services.DeploymentMap {
	return &services.DeploymentMap{}
}

//InitPodMap 初始化Pod存储Map
func (m *K8sMaps) InitPodMap() *services.PodMap {
	return &services.PodMap{}
}

//InitNamespaceMap 初始化Namespace存储Map
func (m *K8sMaps) InitNamespaceMap() *services.NameSpaceMap {
	return &services.NameSpaceMap{}
}

//InitEventMap 初始化Event存储Map
func (m *K8sMaps) InitEventMap() *services.EventMap {
	return &services.EventMap{}
}

//InitIngressMap 初始化Ingress存储Map
func (m *K8sMaps) InitIngressMap() *services.IngressMap {
	return &services.IngressMap{}
}

//InitServiceMap 初始化Service存储Map
func (m *K8sMaps) InitServiceMap() *services.ServiceMap {
	return &services.ServiceMap{}
}

//InitSecretMap 初始化Secret存储Map
func (m *K8sMaps) InitSecretMap() *services.SecretMap {
	return &services.SecretMap{}
}

//InitConfigMapMap 初始化Secret存储Map
func (m *K8sMaps) InitConfigMapMap() *services.ConfigMap {
	return &services.ConfigMap{}
}

//InitNodeMapMap 初始化Node存储Map
func (m *K8sMaps) InitNodeMapMap() *services.NodeMap {
	return &services.NodeMap{}
}

//InitRoleMapMap 初始化Role存储Map
func (m *K8sMaps) InitRoleMapMap() *rbac.RoleMap {
	return &rbac.RoleMap{}
}

//InitRoleBindingMapMap 初始化RoleBinding存储Map
func (m *K8sMaps) InitRoleBindingMapMap() *rbac.RoleBindingMap {
	return &rbac.RoleBindingMap{}
}

//InitServiceAccountMapMap 初始化ServiceAccount存储Map
func (m *K8sMaps) InitServiceAccountMapMap() *rbac.ServiceAccountMap {
	return &rbac.ServiceAccountMap{}
}

//InitClusterRoleMapMap 初始化ClusterRole存储Map
func (m *K8sMaps) InitClusterRoleMapMap() *rbac.ClusterRoleMap {
	return &rbac.ClusterRoleMap{}
}

//InitClusterRoleBindingMapMap 初始化ClusterRoleBinding存储Map
func (m *K8sMaps) InitClusterRoleBindingMapMap() *rbac.ClusterRoleBindingMap {
	return &rbac.ClusterRoleBindingMap{}
}