package cache

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "go/ws/connecter"
    "go/ws/models"
    "strconv"
)

func GetGrabInfo(id int) (GrabInfo *models.Grab, err error) {
    GrabInfo = &models.Grab{}

    conn := connecter.RedisPool.Get()

    defer conn.Close()

    redis.Bytes(conn.Do("select", 1))

    reply, err := redis.Values(conn.Do("hgetall", "QY:GRAB_"+strconv.Itoa(id)))
    if err != nil {
        fmt.Println("hgetall:", err)
        return nil, err
    }
    redis.ScanStruct(reply, &GrabInfo)

    if err != nil {
        return nil, err
    }
    return GrabInfo, nil
}

// 缓存礼物信息
func CacheGiftInfo(giftInfo *models.Grab, id int) {
    conn := connecter.RedisPool.Get()
    var err error
    defer conn.Close()
    redis.Bytes(conn.Do("select", 1))

    if id != 0 {
        _, err = conn.Do("hmset",
            redis.Args{}.Add("QY:GRAB_"+strconv.Itoa(id)).AddFlat(giftInfo)...)
    } else {
        _, err = conn.Do("hmset",
            redis.Args{}.Add("QY:GRAB_"+strconv.Itoa(giftInfo.Id)).AddFlat(giftInfo)...)
    }
    if err != nil {
        fmt.Println(err)
    }
}
