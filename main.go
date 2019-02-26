package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"micro-service/handler"
	"micro-service/subscriber"

	example "micro-service/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro-service.srv.micro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro-service.srv.micro", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro-service.srv.micro", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}