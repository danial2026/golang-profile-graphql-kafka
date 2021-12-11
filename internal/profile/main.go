package main

import (
	context "context"

	profile "github.com/danial2026/golang-profile-graphql-kafka/internal/profile/grpc"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/ports"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/server"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/service"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	application := service.NewApplication(ctx)

	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		profile.RegisterGrpcServerServer(server, svc)
	})
}
