package databases

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

// redispool ..
var redispool *redis.Pool

// GetRedis .
func GetRedis() redis.Conn {
	return redispool.Get()
}

// InitRedis redis初始化
func InitRedis() {
	url := viper.GetString("Redis.url")
	password := viper.GetString("Redis.password")
	maxIdle := viper.GetInt("Redis.maxIdle")
	index := viper.GetInt("Redis.index")
	if url == "" || password == "" {
		log.Fatal("Redis 配置不齐全")
	}

	if maxIdle <= 10 {
		maxIdle = 30
	}
	pool := NewRedisPool(url, password, maxIdle, time.Second*240, index)
loop:
	_, err := pool.Dial()
	if err != nil {
		log.Println("redis error: ", err)
		time.Sleep(3 * time.Second)
		goto loop
	}
	// 清空redis库
	if viper.GetBool("Redis.clearUp") {
		conn := pool.Get()
		conn.Do("flushdb")
		conn.Close()
	}
	redispool = pool
}

// NewRedisPool returns a new Redis connection pool.
func NewRedisPool(redisURL, password string, maxIdle int, idleTimeout time.Duration, databaseIndex int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisURL,
				redis.DialPassword(password),
				redis.DialConnectTimeout(time.Duration(15000)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(15000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(15000)*time.Millisecond))
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			if databaseIndex == 0 {
				databaseIndex = 1
			}
			if _, err := c.Do("SELECT", databaseIndex); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			// if time.Now().Sub(t) < time.Minute {
			// 	return nil
			// }
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}
