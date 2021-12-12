package main

import (
	context "context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/ports"
	postgrpc "github.com/danial2026/golang-profile-graphql-kafka/internal/feed/proto"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/server"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/service"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	application := service.NewApplication(ctx)

	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		postgrpc.RegisterPostsServer(server, svc)
	})
}
