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
	Name       string
	Images     string
	NodeName   string
	CreateTime string
}
