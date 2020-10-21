package config

import (
	"context"
	"employees/controller"
	"github.com/go-redis/redis/v8"
	"log"
)

func ConnectToRedis() {
	// Cache Config
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})

	pong, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal("Couldn't connect to the redis cache", err)
	} else {
		log.Println("Connected to Redis Cache! Response from Redis Ping", pong)
	}
	controller.EmployeeCache(redisClient)
	return
}
