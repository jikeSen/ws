package services

// 赠送礼物
func SendGift(sender, receiver, grabId, nums int) (msg string, err error) {
    // var lock sync.RWMutex

    // 获取礼物信息
    grab, err := GetGiftInfo(grabId)
    if err != nil || grab.GrabPrice == 0 {
        return "礼物不存在", err
    }

    // 读写锁

    return nil, nil
}
