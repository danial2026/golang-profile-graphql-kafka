package main

import (
	context "context"
	"fmt"
	"os"
	"strings"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/gRPC/profile"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/gRPC/server"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/ports"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/service"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	application := service.NewApplication(ctx)

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "grpc":
		server.RunGRPCServer(func(server *grpc.Server) {
			svc := ports.NewGrpcServer(application)
			profile.RegisterGrpcServerServer(server, svc)
		})
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
