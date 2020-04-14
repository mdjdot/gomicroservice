package main

import (
	"log"
	"mservices/services"
	"mservices/student"

	"github.com/micro/go-micro"
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
		micro.Name("student_service"),
		micro.Version("v1.0.0"),
	)

	service.Init()

	student.RegisterStudentServiceHandler(service.Server(), new(services.StudentManager))

	log.Fatal(service.Run())
}
