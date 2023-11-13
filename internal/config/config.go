package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DB    *DBConfig
	Redis *RedisConfig
	Host  string
	Port  string
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

type RedisConfig struct {
	Host        string
	Port        int
	Password    string
	MaxIdle     int `default:"10"`
	MaxActive   int `default:"10"`
	IdleTimeout time.Duration
}

func (r *RedisConfig) Address() string {
	return r.Host + ":" + strconv.Itoa(r.Port)
}

// GetConfig returns config
func GetConfig() *Config {
	db_port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	redis_port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	idle_timeout, _ := time.ParseDuration(os.Getenv("REDIS_IDLE_TIMEOUT"))
	max_idle, _ := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
	max_active, _ := strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE"))
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     os.Getenv("DB_HOST"),
			Port:     db_port,
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Charset:  "utf8",
		},
		Redis: &RedisConfig{
			Host:        os.Getenv("REDIS_HOST"),
			Password:    os.Getenv("REDIS_PASSWORD"),
			Port:        redis_port,
			IdleTimeout: idle_timeout * time.Second,
			MaxIdle:     max_idle,
			MaxActive:   max_active,
		},
		Host: os.Getenv("HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}
}
