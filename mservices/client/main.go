package main

import (
	"context"
	"fmt"
	"mservices/student"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	// "github.com/micro/go-micro/registry"
	// // "github.com/micro/go-plugins/registry/consul"
	// "github.com/micro/go-micro/registry/consul"
)

func main() {
	// reg := consul.NewRegistry(func(op *registry.Options) {
	// 	op.Addrs = []string{
	// 		"127.0.0.1:8500",
	// 	}
	// })
	service := micro.NewService(
		// micro.Registry(reg),
		micro.Name("student.client"),
	)

	service.Init()
	callOption := func(o *client.CallOptions) {
		o.RequestTimeout = 360 * time.Second
		o.DialTimeout = 360 * time.Second
		o.Retries = 3
	}

	studentService := student.NewStudentServiceClient("student_service", service.Client())
	resp, err := studentService.GetStudent(context.Background(), &student.StudentRequest{Name: "davie"}, callOption)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", *resp)

}
