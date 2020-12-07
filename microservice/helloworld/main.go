package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"helloworld/handler"
	"helloworld/subscriber"

	helloworld "helloworld/proto/helloworld"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.helloworld"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.helloworld", service.Server(), new(subscriber.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
