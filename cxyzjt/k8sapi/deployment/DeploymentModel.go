package deployment

type Deployment struct {
	Name       string
	NameSpace  string
	Replicas   [3]int32
	Images     string
	Pods       []Pod
	CreateTime string
}

type Pod struct {
	Name string
	Images string
	NodeName string
	Phase string  // pod 当前所处的阶段
	Message string
	CreateTime string
}
