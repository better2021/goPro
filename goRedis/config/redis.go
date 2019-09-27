package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ConnClient(){
	client := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	pong,err:= client.Ping().Result()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(pong,"redis已启动")

	err = client.Set("myKey","qwerty",0).Err()
	if err!=nil{
		panic(err)
	}

	val,err := client.Get("myKey").Result()

	if err == redis.Nil{
		fmt.Println("key is does not exist")
	}else if err !=nil {
		fmt.Println(err,"--")
	}else {
		fmt.Println("val",val)
	}

}