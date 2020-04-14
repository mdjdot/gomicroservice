package main

import (
	"encoding/json"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/service"
	"github.com/micro/go-plugins/broker/mqtt"
)

func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv"),
		micro.Version("1.2.0"),
		micro.Broker(mqtt.NewBroker()),
	)

	pubSub := service.Server().Options().Broker
	_, err := pubSub.Subscribe("go.micro.srv.message", func(event broker.Event) error {
		var req *message.StudentRequest
		if err := json.Unmarshal(event.Message().Body,&req);err!=nil{
			rerurn err
		}
		fmt.Println("接收到消息：",req)
		return nil
	})
}
