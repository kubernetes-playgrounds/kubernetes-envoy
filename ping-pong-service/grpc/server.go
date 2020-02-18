package pingPongServer

import (
	"context"

	pingpong "github.com/kubernetes-playgrounds/kubernetes-envoy/ping-pong-service/proto"
)

type PingPongServer struct {
}

func NewServer() pingpong.PingPongServer {
	return &PingPongServer{}
}

func (s *PingPongServer) SendPing(context.Context, *pingpong.Ping) (*pingpong.Pong, error) {
	return &pingpong.Pong{Message: "pong"}, nil
}
