package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisCli *redis.Client

func init() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	result, err := RedisCli.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("PING:%s\n", result)
}
