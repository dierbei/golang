package wscore

import "github.com/gorilla/websocket"

type WsShellClient struct {
	client *websocket.Conn
}

func NewWsShellClient(client *websocket.Conn) *WsShellClient {
	return &WsShellClient{client: client}
}

func (ws *WsShellClient) Write(p []byte) (n int, err error) {
	err = ws.client.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (ws *WsShellClient) Read(p []byte) (n int, err error) {
	_, b, err := ws.client.ReadMessage()
	if err != nil {
		return 0, err
	}
	return copy(p, string(b)), nil
}
