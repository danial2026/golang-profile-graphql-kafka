package ports

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/app"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) CreateAccount(ctx context.Context, user domain.User) (*empty.Empty, error) {

	if err := g.app.Commands.CreateAccount.Handle(ctx, user); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) Follow(ctx context.Context, username1 string, username2 string) (*empty.Empty, error) {

	if err := g.app.Commands.Follow.Handle(ctx, username1, username2); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) Unfollow(ctx context.Context, username1 string, username2 string) (*empty.Empty, error) {

	if err := g.app.Commands.Unfollow.Handle(ctx, username1, username2); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
