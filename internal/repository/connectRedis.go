package repository

import (
	"fmt"
	"time"

	"github.com/coderj001/smart_city_traffic_management_system/internal/config"
	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

// Initialize the Redis connection
func ConnectRedis() error {
	redisConfig := config.GetConfig().Redis
	RedisConn = &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		MaxActive:   redisConfig.MaxActive,
		IdleTimeout: redisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConfig.Address())
			if err != nil {
				fmt.Println("Failed to connect to Redis:", err)
				return nil, err
			}
			if redisConfig.Password != "" {
				if _, err := c.Do("AUTH", redisConfig.Password); err != nil {
					c.Close()
					fmt.Println("Failed to authenticate with Redis:", err)
					return nil, err
				}
			}
      fmt.Println("🚀 Connected to Redis successfully")
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}
