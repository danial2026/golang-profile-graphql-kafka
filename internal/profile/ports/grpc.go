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
	_ proto.UnimplementedUserServer
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

func ConvertCreateAccountRequestToDomain(user proto.CreateAccountRequest) domain.User {
	return domain.User{
		ID:        user.Account.GetID(),
		Email:     user.Account.GetEmail(),
		Username:  user.Account.GetUsername(),
		Followers: ConvertListCreateAccountRequestToDomain(user.Account.GetFollowers()),
		Following: ConvertListCreateAccountRequestToDomain(user.Account.GetFollowing()),
	}
}

func ConvertListCreateAccountRequestToDomain(users []*proto.Account) []domain.User {

	dominusers := []domain.User{}

	for _, user := range users {

		dominusers = append(dominusers, ConvertAcountToDomain(*user))
	}

	return dominusers
}

func ConvertAcountToDomain(user proto.Account) domain.User {
	return domain.User{
		ID:        user.GetID(),
		Email:     user.GetEmail(),
		Username:  user.GetUsername(),
		Followers: ConvertListCreateAccountRequestToDomain(user.GetFollowers()),
		Following: ConvertListCreateAccountRequestToDomain(user.GetFollowing()),
	}

}
