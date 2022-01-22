package wscore

import (
	"time"

	"github.com/gorilla/websocket"
)

type WsClient struct {
	conn      *websocket.Conn
	readChan  chan *WsMessage //读队列 (chan)
	closeChan chan byte       // 失败队列
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn, readChan: make(chan *WsMessage), closeChan: make(chan byte)}
}
func (this *WsClient) Ping(waittime time.Duration) {
	for {
		time.Sleep(waittime)
		err := this.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			ClientMap.Remove(this.conn)
			return
		}
	}
}
func (this *WsClient) ReadLoop() {
	for {
		t, data, err := this.conn.ReadMessage()
		if err != nil {
			this.conn.Close()
			ClientMap.Remove(this.conn)
			this.closeChan <- 1
			break
		}
		this.readChan <- NewWsMessage(t, data)
	}
}
