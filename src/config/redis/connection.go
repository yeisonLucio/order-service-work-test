package redis

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Connect() {
	redisDB := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(redisDB)
	if err != nil {
		log.Default().Println("connection redis error: ", err.Error())
		return
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	fmt.Println("#### redis connected ####")
}
