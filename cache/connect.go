package cache

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

// 连接redis数据库
func ConnctRedis() {
	dsn := os.Getenv("REDIS_DSN")
	// 连接redis
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		log.Fatal("连接REDIS数据库失败")
	} else {
		fmt.Println("")
		fmt.Println("\033[0;32m------------REDIS连接成功------------\033[0m")
	}
	RDB = redis.NewClient(opt)
}
