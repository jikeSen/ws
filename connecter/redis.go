package connecter

import (
    "github.com/gomodule/redigo/redis"
    "go/ws/library"
    "time"
)

var RedisPool *redis.Pool

func ConnectRedis() error {
    RedisPool = &redis.Pool{
        MaxIdle:     35,
        MaxActive:   35,
        IdleTimeout: 200,
        Dial: func() (redis.Conn, error) {
            conn, err := redis.Dial("tcp", library.RedisSet.Host+ ":" +library.RedisSet.Port)
            if err != nil {
                return nil, err
            }
            return conn, err
        },
        TestOnBorrow: func(conn redis.Conn, t time.Time) error {
            _, err := conn.Do("PING")
            return err
        },
    }

    return nil
}
