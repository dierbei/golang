package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	pool "github.com/meitu/go-redis-pool"
	"log"
	"time"
)

var (
	redisClient *redis.Client
)

func main() {

	redisPool, err := pool.NewHA(&pool.HAConfig{
		Master: "192.168.75.130:6379",
		Slaves: []string{
			"192.168.75.129:6379",
			"192.168.75.131:6379",
		},
		Password: "123456",
		PollType: pool.PollByRoundRobin,
	})

	if err != nil {
		log.Println("pool connect failed")
	}

	fmt.Println(redisPool.Set("xlt", "aaa", 1000000000000).Err())

	go func() {
		for true {
			fmt.Println(redisPool.Get("xlt").String(), " get操作")
			time.Sleep(3 * time.Second)
		}
	}()

	for i := 0; i < 1000; i++ {
		fmt.Println(redisPool.Set(fmt.Sprintf("name%d", i), "aaa", 1000000000).Err(), " set操作")
		time.Sleep(10 * time.Second)
	}

	//ExampleNewFailoverClient()
	//fmt.Println(redisClient.Set(context.Background(), "xlt", "aaa", 1000000000000).Err())
	//
	//go func() {
	//	for true {
	//		fmt.Println(redisClient.Get(context.Background(), "xlt").String(), " get操作")
	//		time.Sleep(3 * time.Second)
	//	}
	//}()
	//
	//for i := 0; i < 1000; i++ {
	//	fmt.Println(redisClient.Set(context.Background(), fmt.Sprintf("name%d", i), "aaa", 100000000000).Err(), " set操作")
	//	time.Sleep(10 * time.Second)
	//}
}

func ExampleNewFailoverClient() {
	// See http://redis.io/topics/sentinel for instructions how to
	// setup Redis Sentinel.
	redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		Password: "123456",
		MasterName:    "mymaster",
		SentinelAddrs: []string{"192.168.75.129:26379", "192.168.75.130:26379", "192.168.75.131:26379"},
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Print("redisClient.Ping failed ", err.Error())
	}
}