package chat

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/websocket"
    "log"
    "runtime/debug"
    "time"
)

// 用户连接
type Client struct {
    Addr        string          // 客户端地址
    Conn        *websocket.Conn // 用户连接
    Uid         int             // 用户id
    DataPackage chan []byte     // 待发送的数据
    ConnectTime int             // 连接时间
}

func CreateClient(conn *websocket.Conn, addr string, uid int) (client *Client) {
    client = &Client{
        Addr:        addr,
        Conn:        conn,
        Uid:         uid,
        DataPackage: make(chan []byte, 100),
        ConnectTime: int(time.Now().Unix()),
    }
    return
}

// 读取消息 将消息加入
func (client *Client) RedMessage() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("write stop", string(debug.Stack()), r)
        }
    }()

    defer func() {
        close(client.DataPackage)
    }()

    for {
        _, msg, err := client.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        if err != nil {
            return
        }
        HandleGetMsg(client,msg)
    }
}

//发送消息 读取客户端package 发送
func (client *Client) WriteMessage() {

    for {
        select {
        case data, ok := <-client.DataPackage:
            if ok {
                sendErr := client.Conn.WriteMessage(1, data)
                if sendErr != nil {
                    fmt.Println("数据发送完成：%s，%s", client.Addr, data)
                }
            }
        }
    }
}

// client 数据包写入数据，通过通信来共享内存
func (client *Client) SendMsg(code int, msg, data string) (err error) {

    var Msg = &RespMsg{
        Code: code,
        Msg:  msg,
        Data: data,
    }

    packageByte, err := json.Marshal(Msg)

    if err != nil {
        return err
    }

    client.DataPackage <- packageByte
    return
}
