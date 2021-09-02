package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	second time.Duration = 1000000000
)

func main() {
	ExampleNewFailoverClient()
}

func ExampleNewFailoverClient() {
	// See http://redis.io/topics/sentinel for instructions how to
	// setup Redis Sentinel.
	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		Password:      "123456",
		MasterName:    "mymaster",
		SentinelAddrs: []string{"192.168.75.129:26379", "192.168.75.130:26379", "192.168.75.131:26379"},
		RouteRandomly: true,
	})
	rdb.Ping(context.Background())

	rdb.Set(context.Background(), "xlt", "xiaolatiao", second*1000)

	go func() {
		for i := 0; i < 1000; i++ {
			rdb.Set(context.Background(), fmt.Sprintf("name%d", i), "xiaolatiao", second*60)
			time.Sleep(time.Second * 10)
		}
	}()

	for {
		fmt.Println(rdb.Get(context.Background(), "xlt"))
		time.Sleep(time.Second * 2)
	}
}
