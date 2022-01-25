package models

type IngressModel struct {
	Name       string
	NameSpace  string
	CreateTime string
	Host       string
	Options    IngressOptions
}

type IngressOptions struct {
	IsCros    bool
	IsRewrite bool
}

type IngressPath struct {
	Path    string `json:"path"`
	SvcName string `json:"svc_name"`
	Port    string `json:"port"`
}

type IngressRules struct {
	Host  string         `json:"host"`
	Paths []*IngressPath `json:"paths"`
}

type IngressPost struct {
	Name        string
	Namespace   string
	Rules       []*IngressRules
	Annotations string
}
