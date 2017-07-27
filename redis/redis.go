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
	client := NewClient()
	value,err :=client.Get(key).Result()
	defer client.Close()
	if err != nil {
		return ""
	}
	return value
}

func SetValueToRedis(key string, value interface{},exd int)  {
	client := NewClient()
	err := client.Set(key,value,time.Duration(exd)).Err()
	defer client.Close()
	if err != nil {
		panic(err)
	}
}