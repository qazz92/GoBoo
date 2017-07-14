package redis

import (
	"github.com/go-redis/redis"
	"time"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewClient() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("INMISHOST")

	client := redis.NewClient(&redis.Options{
		Addr:     url+":6379",
		Password: "Q!2dltnals", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func GetValueFromRedis(key string) string {
	value,err := NewClient().Get(key).Result()

	if err != nil {
		return ""
	}

	return value
}

func SetValueToRedis(key string, value interface{},exd int)  {
	err := NewClient().Set(key,value,time.Duration(exd)).Err()
	if err != nil {
		panic(err)
	}
}