package client

import (
	"log"

	proto "github.com/danial2026/golang-profile-graphql-kafka/internal/common/proto/profile"
	"google.golang.org/grpc"
)

var UserClient proto.UserClient

func init() {
	UserClient = proto.NewUserClient(makeConnection(":50052"))
}

func makeConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	return conn
}
