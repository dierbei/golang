package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"math/rand"
	"time"
)

var (
	masterUrl = "root:root@tcp(192.168.75.129:3307)/xlt2?charset=utf8mb4&parseTime=True&loc=Local"
	slave1Url = "root:root@tcp(192.168.75.129:3308)/xlt2?charset=utf8mb4&parseTime=True&loc=Local"
	slave2Url = "root:root@tcp(192.168.75.129:3309)/xlt2?charset=utf8mb4&parseTime=True&loc=Local"
)

var DB *gorm.DB
var err error

func main() {
	initMysql()

	//for i := 0; i < 5; i++ {
	//	user := User{UserName: fmt.Sprintf("name%d", i)}
	//	DB.Create(&user)
	//	fmt.Println("插入成功")
	//	time.Sleep(time.Second * 20)
	//}

	u1 := make([]User, 0)
	//u2 := make([]User, 0)
	for i := 0; i < 50; i++ {
		DB.Find(&u1)
		//DB.Find(&u2)
		for _, v := range u1 {
			fmt.Println(v)
		}
		fmt.Println("查询成功u1")
		//for _, v := range u1 {
		//	fmt.Println(v)
		//}
		//fmt.Println("查询成功u2")
		time.Sleep(time.Second * 1)
	}
}

func initMysql() {
	DB, err = gorm.Open(mysql.Open(masterUrl), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err.Error())
		return
	}

	//DB.AutoMigrate(&User{})

	err = DB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(slave1Url), mysql.Open(slave2Url)},

		Policy:   dbresolver.RandomPolicy{},
	}, &User{}))
	if err != nil {
		fmt.Println("注册从数据库失败", err.Error())
	}
}

type RoundRobin struct {
}

func (r RoundRobin) Resolve(connPools []gorm.ConnPool) gorm.ConnPool {
	rand.Seed(time.Now().Unix())
	return connPools[rand.Intn(len(connPools))]
}