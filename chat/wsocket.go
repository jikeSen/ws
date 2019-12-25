package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrade = websocket.Upgrader{
	HandshakeTimeout: time.Second * 30,
	ReadBufferSize:   2048,
	WriteBufferSize:  2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 定义客户端
type Client struct {
	conn *websocket.Conn
	Id   string
}

// 定义在线用户
type lineUserList struct {
	List map[string]*Client
}


// 创建新的服务
func NewSocketServer(w http.ResponseWriter, r *http.Request, id string) (cliect *Client, err error) {

	conn, err := upgrade.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	cliect = &Client{
		conn: conn,
		Id:   id,
	}
	return
}

// 发送消息
func (client *Client) SendMsg(msgType int, content string) (err error) {

	var Msg = &Message{
		ID:      client.Id,
		Content: content,
		Type:    msgType,
	}

	sendErr := client.conn.WriteJSON(Msg)

	if sendErr != nil {
		log.Println("Error :", sendErr)
		log.Println("messageBody: ", Msg)
		log.Println("client: ", client)
		client.conn.Close()
	}
	return
}
