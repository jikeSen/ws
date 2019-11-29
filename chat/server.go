package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 启动服务 main 调用
func StartServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hi 小伙子 你很棒"))
	})
	http.HandleFunc("/ws", wsHandle)

	s := &http.Server{
		Addr:           ":8085",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func wsHandle(writer http.ResponseWriter, request *http.Request) {
	if websocket.IsWebSocketUpgrade(request) == false {
		writer.Write([]byte("您请求的不是WebSocket协议"))
	}

	client, err := NewSocketServer(writer, request, createSocketId(7))
	if err != nil {
		fmt.Println(err)
		panic("socket创建错误")
	}

	defer client.conn.Close()

	msgStr := fmt.Sprintf("您好，您的ID：%s",client.Id)

	client.SendMsg(SysMsg,msgStr)

	for {
		messageType, p, err := client.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(messageType)

		if err := client.SendMsg(TxtMsg,string(p)); err != nil {
			log.Println(err)
			return
		}
	}
}

// 创建客户端id
func createSocketId(l int) string {
	var str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
