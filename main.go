package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "go/ws/chat"
    "go/ws/connecter"
    "go/ws/library"
    "log"
    "net/http"
    "time"
)

func init() {
    library.Init()

    mysqlerr := connecter.MysqlConnect()
    if mysqlerr != nil {
        panic(fmt.Errorf("mysql_server err: %s \n", mysqlerr))
    }
    err := connecter.ConnectRedis()
    if err != nil {
        panic(fmt.Errorf("redis_server err: %s \n", err))
    }
}

func main() {
    http.HandleFunc("/", rootTest)

    http.HandleFunc("/redis", testRedis)

    // start ws server
    go func() {
        err := chat.StartWebSocket()
        if err != nil {
            panic(fmt.Errorf("socket 服务启动错误~ %s \n", err))
        }
    }()

    httpPort := library.GetSingleConf("app", "HttpPort")

    HttpServer := &http.Server{
        Addr:           ":" + httpPort.(string),
        Handler:        nil,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(HttpServer.ListenAndServe())
}

func rootTest(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("hi 小伙子 你很棒"))
}

func testRedis(writer http.ResponseWriter, request *http.Request) {
    conn := connecter.RedisPool.Get()

    defer conn.Close()

    redis.Bytes(conn.Do("select", 1))
    reply, err := redis.Bytes(conn.Do("GET", "LK:Flag_SetUidLine_101011"))
    if err != nil {
        panic(err)
    }
    writer.Write([]byte(reply))
}
