package chat

import (
    "fmt"
    "github.com/json-iterator/go"
    "go/ws/library/code"
    "go/ws/services"
)

const (
    SysMsg   = 1 << iota // 系统消息
    LoginMsg             // 登录消息
    ExitMsg              // 退出消息
    RegMsg               // 注册成功
    TxtMsg               // 文本消息
    SoundMsg             // 音频消息
    MediaMsg             // 媒体消息
    C2cMsg               // 单聊消息
    GiftMsg              // 送礼消息
)

// 定义发送的消息
type Message struct {
    ID      string // 消息id
    Content string // 消息内容
    SendAt  int64  // 发消息用户
    Type    int    // 消息类型
    Cmd     string // 执行的action （考虑和type的冲突）暂时保留
}

// 响应消息
type RespMsg struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

// c2c消息
type C2CMsg struct {
    Message
    ToAccount string
}

type GiftMsgs struct {
    CMD    string
    GiftId int
    Sender int
    Getter int
    Num    int
}

// 处理客户端消息
func HandleGetMsg(client *Client, message []byte) {
    fmt.Println("处理数据.....", client.Addr, string(message))

    msgObj := GiftMsgs{}
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    jsonerr := json.Unmarshal(message, &msgObj)
    if jsonerr != nil {
        fmt.Println(jsonerr)
        return
    }

    // 查询 送礼 扣费 流水等业务code

    switch msgObj.CMD {
    case "gift":
        services.SendGift(msgObj.Sender, msgObj.Getter, msgObj.GiftId, msgObj.Num)
        break
    }

    client.SendMsg(code.SUCCESS, code.GetCodeMsg(code.SUCCESS), "测试业务")
}
