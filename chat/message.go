package chat

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
    Content string
    SendAt  int64 // 发消息用户
    Type    int
}

// c2c消息
type C2CMsg struct {
    Message
    ToAccount string
}
