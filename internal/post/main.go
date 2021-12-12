package main

import (
	context "context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/ports"
	postgrpc "github.com/danial2026/golang-profile-graphql-kafka/internal/post/proto"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/server"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/service"

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
