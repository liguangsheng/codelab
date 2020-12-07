package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	helloworld "helloworld/proto/helloworld"
)

type Helloworld struct{}

func (e *Helloworld) Handle(ctx context.Context, msg *helloworld.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *helloworld.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
