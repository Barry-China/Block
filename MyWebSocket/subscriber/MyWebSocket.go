package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	MYWebSocket "vtoken_digiccy_go/MyWebSocket/proto/MyWebSocket"
)

type MyWebSocket struct{}

func (e *MyWebSocket) Handle(ctx context.Context, msg *MYWebSocket.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *MYWebSocket.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
