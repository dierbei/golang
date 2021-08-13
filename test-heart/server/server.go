package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	server := ":7777"

	listen, err := net.Listen("tcp", server)
	if err != nil {
		log.Printf("error [net.Listen], message:%v\n", err)
		return
	}

	log.Println("waiting client connect...")
	for true {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("error [listen.Accept], message:%v\n", err)
			continue
		}

		log.Printf("success [listen.Accept], remoteAddr:%v\n", conn.RemoteAddr())
		go handlerConnection(conn)
	}
}

func handlerConnection(conn net.Conn) {
	buffer := make([]byte, 4096)

	// 监听客户端是否活跃
	isLive := make(chan bool)

	go func() {
		for {
			n, err := conn.Read(buffer)
			if err != nil && err != io.EOF {
				conn.Close()
				return
			}

			// 去除最后的换行符
			data := buffer[:n-1]

			log.Println("接收到消息：", string(data))
			isLive <- true
		}
	}()

	detectCount := 0
	for true {
		select {
		case <-isLive:
			// 当前客户端是活跃的，重置定时器
		case <-time.After(time.Second * 10):
			// 已经超时，主动给客户端发送探测报文
			// 发送次数设定为3次，超过三次断开与用户连接
			detectCount++
			if detectCount % 3 == 0 {
				close(isLive)
				conn.Close()
				log.Printf("close connect, client addr = %v\n", conn.RemoteAddr().String())
				return
			}
			fmt.Printf("第%d次发送探测报文\n", detectCount)
			conn.Write([]byte(fmt.Sprintf("%s 你在吗?.", conn.RemoteAddr().String())))
		}
	}
}