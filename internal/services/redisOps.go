package services

import (
	"encoding/json"

	"github.com/coderj001/smart_city_traffic_management_system/internal/repository"
	"github.com/gomodule/redigo/redis"
)

// Set a key-value pair
func Set(key string, data interface{}, time int) error {
	conn := repository.RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check if a key exists
func Exists(key string) bool {
	conn := repository.RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get a value by key
func Get(key string) ([]byte, error) {
	conn := repository.RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// Delete a value by key
func Delete(key string) (bool, error) {
	conn := repository.RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
