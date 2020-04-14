package main

import (
	"encoding/json"
	"log"

	"github.com/go-delve/delve/service"
	"github.com/micro/go-micro/broker"
)

func main() {
	brok := service.Server().Options().Broker
	if err := brok.Connect(); err != nil {
		log.Fatal(" broker connection failed, error : ", err.Error())
	}

	student := &message.Student{Name: "davie", Classes: "软件工程专业", Grade: 80, Phone: "12345678901"}
	msgBody, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err.Error())
	}
	msg := &broker.Message{
		Header: map[string]string{
			"name": student.Name,
		},
		Body: msgBody,
	}

	err = brok.Publish("go.micro.srv.message", msg)
	if err != nil {
		log.Fatal(" 消息发布失败：%s\n", err.Error())
	} else {
		log.Print("消息发布成功")
	}
}
