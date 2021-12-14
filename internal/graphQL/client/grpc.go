package client

import (
	proto "github.com/danial2026/golang-profile-graphql-kafka/internal/common/proto/profile"
	"google.golang.org/grpc"
)

var UserClient proto.UserClient

func ConnectToService() {
	UserClient = proto.NewUserClient(makeConnection("localhost:50052"))

	// UserClient.CreateAccount()
}

func makeConnection(addr string) *grpc.ClientConn {
	con, err := grpc.Dial(addr)

	if err != nil {
	}
	return con
}
