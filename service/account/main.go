package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro"
	"log"
)

func main(){
	// 创建一个service
	service:=micro.NewService(
		micro.Name("go.micro.service.user"),
	)

	service.Init()

	proto.RegisterUserServiceHandler(service.Server(),new(hanlder.User))
	if err := service.Run();err!=nil{
		log.Println(err)
	}
}

