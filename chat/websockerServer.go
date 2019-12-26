package chat

import (
    "fmt"
    "github.com/gorilla/websocket"
    "go/ws/library"
    "go/ws/library/code"
    "go/ws/models"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

type WsAuth struct {
    Uid   int
    Token string
    App   int
    Tune  string
}

func StartWebSocket() error {
    http.HandleFunc("/ws", wsHandle)

    webSocketPort := library.GetSingleConf("app", "WebsocketPort")

    err := http.ListenAndServe(":"+webSocketPort.(string), nil)

    if err != nil {
        return err
    }
    return nil
}

// 创建socket 服务
func wsHandle(writer http.ResponseWriter, request *http.Request) {
    if websocket.IsWebSocketUpgrade(request) == false {
        writer.Write([]byte(code.GetCodeMsg(code.SOCKET_PROTOCOL_ERR)))
        return
    }

    err := authorizedConnect(writer, request)
    if err != nil {
        writer.Write([]byte(code.GetCodeMsg(code.AUTH_ERR)))
        return
    }

    client, err := NewSocketServer(writer, request, createSocketId(7))
    if err != nil {
        writer.Write([]byte(code.GetCodeMsg(code.SOCKET_CREAT_ERR)))
        return
    }

    defer client.conn.Close()

    client.SendMsg(SysMsg, fmt.Sprintf("您好，您的ID：%s", client.Id, ))

    // 待优化...
    // 拆分 read 和 write 使用携程并行处理

    for {
        _, p, err := client.conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        if err := client.SendMsg(TxtMsg, string(p)); err != nil {
            log.Println(err)
            return
        }
    }
}

func authorizedConnect(writer http.ResponseWriter, request *http.Request) error {
    queryForm, err := url.ParseQuery(request.URL.RawQuery)
    if err == nil && len(queryForm["id"]) > 0 {
        writer.Write([]byte(code.GetCodeMsg(code.INVALID_PARAMS)))
        return err
    }

    userAuth := WsAuth{}

    uid, _ := strconv.Atoi(queryForm.Get("uid"))
    userAuth.Uid = uid
    app, _ := strconv.Atoi(queryForm.Get("app"))
    userAuth.App = app
    userAuth.Token = queryForm.Get("token")
    userAuth.Tune = queryForm.Get("tune")

    // 验签
    user := models.User{}
    uu := user.GetUserInfoByUid(uid)
    fmt.Println(uu.Token)
    return nil
}

//
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
