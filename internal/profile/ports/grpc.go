package ports

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/app"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"
	proto "github.com/danial2026/golang-profile-graphql-kafka/internal/profile/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
	proto.UnimplementedUserServer
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) CreateAccount(ctx context.Context, user *proto.CreateAccountRequest) (*empty.Empty, error) {

	if err := g.app.Commands.CreateAccount.Handle(ctx, ConvertCreateAccountRequestToDomain(*user)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) FollowUser(ctx context.Context, user *proto.FollowUserRequest) (*empty.Empty, error) {

	if err := g.app.Commands.Follow.Handle(ctx, user.GetUsername1(), user.GetUsername2()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) UnfollowUser(ctx context.Context, user *proto.UnfollowUserRequest) (*empty.Empty, error) {

	if err := g.app.Commands.Unfollow.Handle(ctx, user.GetUsername1(), user.GetUsername2()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func ConvertCreateAccountRequestToDomain(user proto.CreateAccountRequest) domain.User {
	return domain.User{
		ID:        user.Account.GetID(),
		Email:     user.Account.GetEmail(),
		Username:  user.Account.GetUsername(),
		Following: []domain.User{},
		Followers: []domain.User{},
	}
}
