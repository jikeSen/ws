package chat

import (
    _ "bufio"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "github.com/gorilla/websocket"
    "go/ws/connecter"
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

// 启动服务 main 调用
func StartServer() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        writer.Write([]byte("hi 小伙子 你很棒"))
    })

    http.HandleFunc("/redis", func(writer http.ResponseWriter, request *http.Request) {
        conn := connecter.RedisPool.Get()

        defer conn.Close()

        redis.Bytes(conn.Do("select", 1))
        reply, err := redis.Bytes(conn.Do("GET", "LK:Flag_SetUidLine_101011"))
        if err != nil {
            panic(err)
        }

        writer.Write([]byte(reply))
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

    // 获取参数
    queryForm, err := url.ParseQuery(request.URL.RawQuery)
    if err == nil && len(queryForm["id"]) > 0 {
        writer.Write([]byte("参数错误"))
    }

    userAuth := WsAuth{}

    uid, _ := strconv.Atoi(queryForm.Get("uid"))
    userAuth.Uid = uid
    app, _ := strconv.Atoi(queryForm.Get("app"))
    userAuth.App = app
    userAuth.Token = queryForm.Get("token")
    userAuth.Tune = queryForm.Get("tune")

    // 获取用户的信息 进行加密验证
    user := models.User{}
    uu := user.GetUserInfoByUid(uid)
    fmt.Println(uu.Token)

    client, err := NewSocketServer(writer, request, createSocketId(7))
    if err != nil {
        fmt.Println(err)
        panic("socket创建错误")
    }

    defer client.conn.Close()

    client.SendMsg(SysMsg, fmt.Sprintf("您好，您的ID：%s", client.Id, ))

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
