package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"time"
)

type MyActor struct {
	count int
}

func (state *MyActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.Response:
		state.count++
		fmt.Println(state.count)
	}
}

func main() {
	remote.Start("localhost:8090")

	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor { return &MyActor{} })
	pid:= context.Spawn(props)
	message := &messages.Echo{Message: "hej", Sender: pid}

	// this is to spawn remote actor we want to communicate with
	spawnResponse, _ := remote.SpawnNamed("localhost:8091", "myactor", "hello", time.Second)

	// get spawned PID
	spawnedPID := spawnResponse.Pid
	for i := 0; i < 10; i++ {
		context.Send(spawnedPID, message)
	}

	console.ReadLine()
}