package code

// 定义错误map
var msgFlag = map[int]string{
    SUCCESS:        "成功",
    ERROR:          "失败",
    INVALID_PARAMS: "请求参数有误",

    AUTH_ERR:            "授权失败，请检测服务配置和授权签名",
    SOCKET_CREAT_ERR:    "服务端创建socket错误，请重试",
    SOCKET_PROTOCOL_ERR: "连接协议错误",
}

// 根据返回的错误码返回定义的错误信息
func GetCodeMsg(code int) (msg string) {
    codeMsg, ok := msgFlag[code]

    if ok {
        return codeMsg
    }
    return msgFlag[ERROR]
}
