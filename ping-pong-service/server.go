package main

import (
	"fmt"
	"log"
	"net"

	pingPongServer "github.com/kubernetes-playgrounds/kubernetes-envoy/ping-pong-service/grpc"
	pingpong "github.com/kubernetes-playgrounds/kubernetes-envoy/ping-pong-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "9091"))
	if err != nil {
		log.Fatalf("Error creating lis: %v", err)
	}
	grpcServer := grpc.NewServer()
	pingPongServ := pingPongServer.NewServer()
	healthServer := health.NewServer()
	healthServer.SetServingStatus("ping-pong", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	pingpong.RegisterPingPongServer(grpcServer, pingPongServ)

	fmt.Printf("Starting grpc server on port:%v", "9091")
	grpcServer.Serve(lis)
}
