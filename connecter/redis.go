package connecter

import (
    "github.com/gomodule/redigo/redis"
    "time"
)

var RedisPool *redis.Pool

func ConnectRedis() error {
    RedisPool = &redis.Pool{
        MaxIdle:     35,
        MaxActive:   35,
        IdleTimeout: 200,
        Dial: func() (redis.Conn, error) {
            conn, err := redis.Dial("tcp", "127.0.0.1:6379")
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
