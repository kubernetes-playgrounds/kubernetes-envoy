package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pingpong "github.com/kubernetes-playgrounds/kubernetes-envoy/ping-pong-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create grpc conn: %v", err)
	}
	cli := pingpong.NewPingPongClient(conn)
	for {
		resp, err := cli.SendPing(context.Background(), &pingpong.Ping{})
		if err != nil {
			log.Fatalf("Received error from server: %v", err)
		}
		fmt.Printf("Received resp from server: %v\n", resp.Message)
		time.Sleep(5 * time.Second)
	}
}
