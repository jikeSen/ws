package main

import (
	"go/ws/chat"
	"go/ws/connecter"
)

func init(){
	connecter.MysqlConnect()
	err := connecter.ConnectRedis()
	if err != nil {
		panic(err)
	}
}

func main(){
	chat.StartServer()
}