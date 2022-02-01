package wscore

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

//外部公共使用,保存所有客户端对象的Map
var ClientMap *ClientMapStruct

func init() {
	ClientMap = &ClientMapStruct{}
}

type ClientMapStruct struct {
	data sync.Map //  key 是客户端IP  value 就是 WsClient连接对象
}

func (this *ClientMapStruct) Store(conn *websocket.Conn) {
	wsClient := NewWsClient(conn)
	this.data.Store(conn.RemoteAddr().String(), wsClient)
	go wsClient.Ping(time.Second * 1)
	go wsClient.ReadLoop() //处理读 循环
	// go wsClient.HandlerLoop() //处理 总控制循环
}

//SendAll 向所有客户端 发送消息--发送deployment列表
func (this *ClientMapStruct) SendAll(v interface{}) {
	this.data.Range(func(key, value interface{}) bool {
		c := value.(*WsClient).conn
		err := c.WriteJSON(v)
		if err != nil {
			this.Remove(c)
			log.Println(err)
		}
		return true
	})
}
func (this *ClientMapStruct) Remove(conn *websocket.Conn) {
	this.data.Delete(conn.RemoteAddr().String())
}
