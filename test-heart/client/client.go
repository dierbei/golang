package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	server := "127.0.0.1:7777"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		log.Printf("error [net.ResolveTCPAddr] message: %v\n", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("error [net.DialTCP] message: %v\n", err)
		return
	}

	log.Println(conn.RemoteAddr().String(), " connect success!")

	send(conn)
	log.Printf("send over")
}

// ReceiveDetectFromServer
// 客户端接收来自服务端的探测报文给予相应
func ReceiveDetectFromServer() {
	// todo
}

func send(conn *net.TCPConn) {
	for i := 0; i < 10 ; i++ {
		words := strconv.Itoa(i) + " Hello I'm Client.\n"
		msg, err := conn.Write([]byte(words))
		if err != nil {
			log.Printf("error [conn.Write] message: %v\n", err)
			return
		}
		log.Printf("发送数据大小：%v, 消息内容: %v", msg, words)
		time.Sleep(time.Second * 3)
	}

	for i := 0; i < 2; i++ {
		time.Sleep(11 * time.Second)
	}
}
