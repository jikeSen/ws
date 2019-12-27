package chat

import (
    "fmt"
    "github.com/gorilla/websocket"
    "go/ws/library"
    "go/ws/library/code"
    "go/ws/models"
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

var upgrade = websocket.Upgrader{
    HandshakeTimeout: time.Second * 30,
    ReadBufferSize:   2048,
    WriteBufferSize:  2048,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
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

    // 连接实例
    conn, err := upgrade.Upgrade(writer, request, nil)

    if err != nil {
        writer.Write([]byte(code.GetCodeMsg(code.SOCKET_CREAT_ERR)))
        return
    }
    queryForm, err := url.ParseQuery(request.URL.RawQuery)
    uid, _ := strconv.Atoi(queryForm.Get("uid"))

    client := CreateClient(conn, conn.RemoteAddr().String(), uid)

    // 协程读写消息
    go client.RedMessage()
    go client.WriteMessage()
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
