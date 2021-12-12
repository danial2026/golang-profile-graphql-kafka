package ports

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/app"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/domain"
	proto "github.com/danial2026/golang-profile-graphql-kafka/internal/post/proto"

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

func (g GrpcServer) CreatePost(ctx context.Context, post *proto.CreatePostRequest) (*empty.Empty, error) {

	if err := g.app.Commands.CreatePost.Handle(ctx, ConvertCreatePostRequestToDomain(*post)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) LikePost(ctx context.Context, likeRequest *proto.LikePostRequest) (*empty.Empty, error) {

	if err := g.app.Commands.LikePost.Handle(ctx, likeRequest.Username, likeRequest.PostId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func ConvertCreatePostRequestToDomain(post proto.CreatePostRequest) domain.Post {
	return domain.Post{
		ID:          post.Post.GetID(),
		ProfileId:   post.Post.GetProfileId(),
		Body:        post.Post.GetBody(),
		CreatedTime: post.Post.GetCreatedTime(),
	}
}
