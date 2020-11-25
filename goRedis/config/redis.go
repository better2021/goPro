package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ConnClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	pong,err:= client.Ping().Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println(pong,"redis已启动")

	return client
}