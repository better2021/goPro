package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goRedis/config"
)

func main(){
	client := config.ConnClient()

	err := client.Set("myKey","qwerty",0).Err()
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

	//gin
	router := gin.Default()
	router.GET("/api/getNums", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"data":val,
		})
	})
	router.Run(":8081")
}

