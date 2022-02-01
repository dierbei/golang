package models

type ConfigMapModel struct {
	Name       string
	NameSpace  string
	CreateTime string
	Data       map[string]string
}

type PostConfigMapModel struct {
	Name      string
	NameSpace string
	Data      map[string]string
	IsUpdate  bool
}
