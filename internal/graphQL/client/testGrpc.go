package client

import (
	"log"

	profileProto "github.com/danial2026/golang-profile-graphql-kafka/internal/common/proto/profile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func testGrpc() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := profileProto.NewUserClient(conn)
	response, err := c.CreateAccount(
		context.Background(), &profileProto.CreateAccountRequest{
			Account: &profileProto.Account{
				Email:     "foo",
				Username:  "fee",
			},
		},
	)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)
}
