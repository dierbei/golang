package main

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

func (model *User) TableName() string {
	return "user"
}
