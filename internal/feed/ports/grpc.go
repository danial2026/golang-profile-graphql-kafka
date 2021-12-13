package ports

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/app"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/domain"
	proto "github.com/danial2026/golang-profile-graphql-kafka/internal/feed/proto"

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

func (g GrpcServer) CreateNotification(ctx context.Context, feed *proto.CreateNotificationRequest) (*empty.Empty, error) {

	if err := g.app.Commands.CreateNotification.Handle(ctx, ConvertCreateFeedRequestToDomain(*feed)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func ConvertCreateFeedRequestToDomain(feed proto.CreateNotificationRequest) domain.Feed {
	return domain.Feed{
		ID:          feed.Feed.GetID(),
		Username:    feed.Feed.GetUsername(),
		PostId: 	 feed.Feed.GetCreatedTime(),
		CreatedTime: feed.Feed.GetCreatedTime(),
	}
}
