package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type RbacCtl struct {
	RoleService           *RoleService           `inject:"-"`
	Client                *kubernetes.Clientset  `inject:"-"`
	ServiceAccountService *ServiceAccountService `inject:"-"`
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

func (ctl *RbacCtl) SaList(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": ctl.ServiceAccountService.ListSa(ns),
	}
}

func (ctl *RbacCtl) ClusterRoles(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ctl.RoleService.ListClusterRoles(),
	}
}

func (ctl *RbacCtl) DeleteClusterRole(c *gin.Context) goft.Json {
	name := c.DefaultQuery("name", "")
	err := ctl.Client.RbacV1().ClusterRoles().Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) CreateClusterRole(c *gin.Context) goft.Json {
	clusterRole := rbacv1.ClusterRole{} //原生的k8s role 对象
	goft.Error(c.ShouldBindJSON(&clusterRole))
	clusterRole.APIVersion = "rbac.authorization.k8s.io/v1"
	clusterRole.Kind = "ClusterRole"
	_, err := ctl.Client.RbacV1().ClusterRoles().Create(c, &clusterRole, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) ClusterRolesDetail(c *gin.Context) goft.Json {
	rname := c.Param("cname") //集群角色名
	return gin.H{
		"code": 20000,
		"data": ctl.RoleService.GetClusterRole(rname),
	}
}

func (ctl *RbacCtl) UpdateClusterRolesDetail(c *gin.Context) goft.Json {
	cname := c.Param("cname") //集群角色名
	clusterRole := ctl.RoleService.GetClusterRole(cname)
	postRole := rbacv1.ClusterRole{}
	goft.Error(c.ShouldBindJSON(&postRole)) //获取提交过来的对象

	clusterRole.Rules = postRole.Rules //目前修改只允许修改 rules，其他不允许。大家可以自行扩展，如标签也允许修改
	_, err := ctl.Client.RbacV1().ClusterRoles().Update(c, clusterRole, metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) UpdateRolesDetail(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	rname := c.Param("rolename")
	role := ctl.RoleService.GetRole(ns, rname)
	postRole := rbacv1.Role{}
	goft.Error(c.ShouldBindJSON(&postRole)) //获取提交过来的对象

	role.Rules = postRole.Rules //目前修改只允许修改 rules，其他不允许。大家可以自行扩展，如标签也允许修改
	_, err := ctl.Client.RbacV1().Roles(role.Namespace).Update(c, role, metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) ClusterRoleBindingList(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": ctl.RoleService.ListClusterRoleBindings(),
	}
}

func (ctl *RbacCtl) CreateClusterRoleBinding(c *gin.Context) goft.Json {
	rb := &rbacv1.ClusterRoleBinding{}
	goft.Error(c.ShouldBindJSON(rb))
	_, err := ctl.Client.RbacV1().ClusterRoleBindings().Create(c, rb, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) AddUserToClusterRoleBinding(c *gin.Context) goft.Json {
	name := c.DefaultQuery("name", "") //clusterrolebinding 名称
	t := c.DefaultQuery("type", "")    //如果没传值就是增加，传值（不管什么代表删除)
	subject := rbacv1.Subject{}        // 传过来
	goft.Error(c.ShouldBindJSON(&subject))
	if subject.Kind == "ServiceAccount" {
		subject.APIGroup = ""
	}
	rb := ctl.RoleService.GetClusterRoleBinding(name) //通过名称获取 clusterrolebinding对象
	if t != "" {                                      //代表删除
		for i, sub := range rb.Subjects {
			if sub.Kind == subject.Kind && sub.Name == subject.Name {
				rb.Subjects = append(rb.Subjects[:i], rb.Subjects[i+1:]...)
				break //确保只删一个（哪怕有同名同kind用户)
			}
		}
	} else {
		rb.Subjects = append(rb.Subjects, subject)
	}
	_, err := ctl.Client.RbacV1().ClusterRoleBindings().Update(c, rb, metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (ctl *RbacCtl) DeleteClusterRoleBinding(c *gin.Context) goft.Json {
	name := c.DefaultQuery("name", "")
	err := ctl.Client.RbacV1().ClusterRoleBindings().Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

//UserAccount 列表
func (ctl *RbacCtl) UaList(c *gin.Context) goft.Json {
	uaPath := "./k8susers" //写死的路径存储证书
	keyReg := regexp.MustCompile(".*_key.pem")
	users := []*UAModel{}
	suffix := ".pem"
	err := filepath.Walk(uaPath, func(p string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		if path.Ext(f.Name()) == suffix {
			if !keyReg.MatchString(f.Name()) {
				users = append(users, &UAModel{
					Name:       strings.Replace(f.Name(), suffix, "", -1),
					CreateTime: f.ModTime().Format("2006-01-02 15:04:05"),
				})
			}
		}
		return nil
	})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": users,
	}

}
//
//func (ctl *RbacCtl) PostUa(c *gin.Context) goft.Json {
//	postModel := &PostUAModel{}
//	goft.Error(c.ShouldBindJSON(postModel))
//	helpers.GenK8sUser(postModel.CN, postModel.O)
//	return gin.H{
//		"code": 20000,
//		"data": "success",
//	}
//
//}
//
//func (ctl *RbacCtl) DeleteUa(c *gin.Context) goft.Json {
//	postModel := &PostUAModel{}
//	goft.Error(c.ShouldBindJSON(postModel))
//	helpers.DeleteK8sUser(postModel.CN)
//	return gin.H{
//		"code": 20000,
//		"data": "success",
//	}
//
//}

func (*RbacCtl) Name() string {
	return "RBACCtl"
}

func (ctl *RbacCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/roles", ctl.Roles)
	goft.Handle("POST", "/roles", ctl.CreateRole)
	goft.Handle("DELETE", "/roles", ctl.DeleteRole)
	goft.Handle("GET", "/roles/:ns/:rolename", ctl.RolesDetail)
	goft.Handle("POST", "/roles/:ns/:rolename", ctl.UpdateRolesDetail)

	goft.Handle("GET", "/rolebindings", ctl.RoleBindingList)
	goft.Handle("POST", "/rolebindings", ctl.CreateRoleBinding)
	goft.Handle("DELETE", "/rolebindings", ctl.DeleteRoleBinding)
	goft.Handle("PUT", "/rolebindings", ctl.AddUserToRoleBinding)

	goft.Handle("GET", "/sa", ctl.SaList)

	goft.Handle("GET", "/clusterroles", ctl.ClusterRoles)
	goft.Handle("DELETE", "/clusterroles", ctl.DeleteClusterRole)
	goft.Handle("POST", "/clusterroles", ctl.CreateClusterRole)
	goft.Handle("GET", "/clusterroles/:cname", ctl.ClusterRolesDetail)
	goft.Handle("POST", "/clusterroles/:cname", ctl.UpdateClusterRolesDetail)

	goft.Handle("GET", "/clusterrolebindings", ctl.ClusterRoleBindingList)
	goft.Handle("POST", "/clusterrolebindings", ctl.CreateClusterRoleBinding)
	goft.Handle("PUT", "/clusterrolebindings", ctl.AddUserToClusterRoleBinding)
	goft.Handle("DELETE", "/clusterrolebindings", ctl.DeleteClusterRoleBinding)

	goft.Handle("GET", "/ua", ctl.UaList)
	//goft.Handle("POST", "/ua", ctl.PostUa)
	//goft.Handle("DELETE", "/ua", ctl.DeleteUa)
}
