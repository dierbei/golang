package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type RbacCtl struct {
	RoleService *RoleService          `inject:"-"`
	Client      *kubernetes.Clientset `inject:"-"`
}

func NewRBACCtl() *RbacCtl {
	return &RbacCtl{}
}

func (ctl *RbacCtl) Roles(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.RoleService.ListRoles(ns),
	}
}

func (ctl *RbacCtl) RoleBindingList(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.RoleService.ListRoleBindings(ns),
	}
}
func (ctl *RbacCtl) CreateRoleBinding(c *gin.Context) goft.Json {
	rb := &rbacv1.RoleBinding{}
	goft.Error(c.ShouldBindJSON(rb))

	for i, v := range rb.Subjects {
		if v.Kind == "ServiceAccount" {
			rb.Subjects[i].APIGroup = ""
		}
	}

	_, err := ctl.Client.RbacV1().RoleBindings(rb.Namespace).Create(c, rb, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

//AddUserToRoleBinding 从rolebinding中 增加或删除用户
func (ctl *RbacCtl) AddUserToRoleBinding(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "") //rolebinding 名称
	t := c.DefaultQuery("type", "")    //如果没传值就是增加，传值（不管什么代表删除)
	subject := rbacv1.Subject{}        // 传过来
	goft.Error(c.ShouldBindJSON(&subject))
	if subject.Kind == "ServiceAccount" {
		subject.APIGroup = ""
	}
	rb := ctl.RoleService.GetRoleBinding(ns, name) //通过名称获取 rolebinding对象
	if t != "" {                                   //代表删除

		for i, sub := range rb.Subjects {
			if sub.Kind == subject.Kind && sub.Name == subject.Name {
				rb.Subjects = append(rb.Subjects[:i], rb.Subjects[i+1:]...)
				break //确保只删一个（哪怕有同名同kind用户)
			}
		}
	} else {
		rb.Subjects = append(rb.Subjects, subject)
	}
	_, err := ctl.Client.RbacV1().RoleBindings(ns).Update(c, rb, metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) DeleteRole(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "")
	err := ctl.Client.RbacV1().Roles(ns).Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) DeleteRoleBinding(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "")
	err := ctl.Client.RbacV1().RoleBindings(ns).Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) CreateRole(c *gin.Context) goft.Json {
	role := rbacv1.Role{} //原生的k8s role 对象
	goft.Error(c.ShouldBindJSON(&role))
	role.APIVersion = "rbac.authorization.k8s.io/v1"
	role.Kind = "Role"
	_, err := ctl.Client.RbacV1().Roles(role.Namespace).Create(c, &role, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) RolesDetail(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	rname := c.Param("rolename")
	return gin.H{
		"code": 20000,
		"data": ctl.RoleService.GetRole(ns, rname),
	}
}

func (*RbacCtl) Name() string {
	return "RBACCtl"
}

func (ctl *RbacCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/roles", ctl.Roles)
	goft.Handle("POST", "/roles", ctl.CreateRole)
	goft.Handle("DELETE", "/roles", ctl.DeleteRole)
	goft.Handle("GET", "/rolebindings", ctl.RoleBindingList)
	goft.Handle("POST", "/rolebindings", ctl.CreateRoleBinding)
	goft.Handle("DELETE", "/rolebindings", ctl.DeleteRoleBinding)
	goft.Handle("PUT", "/rolebindings", ctl.AddUserToRoleBinding)
	goft.Handle("GET", "/roles/:ns/:rolename", ctl.RolesDetail)
}
