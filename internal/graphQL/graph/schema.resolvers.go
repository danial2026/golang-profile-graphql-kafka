package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	profileProto "github.com/danial2026/golang-profile-graphql-kafka/internal/common/proto/profile"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/graphql/client"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/graphql/graph/generated"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*bool, error) {
	client.UserClient.CreateAccount(ctx, ConvertToCreateAccountRequest(input))

	log.Printf("%s", ConvertToCreateAccountRequest(input))
	log.Printf("%s", input)

	response := true

	return &response, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func ConvertToCreateAccountRequest(user model.NewUser) *profileProto.CreateAccountRequest {
	return &profileProto.CreateAccountRequest{
		Account: &profileProto.Account{
			Email:    user.Email,
			Username: user.Username,
		},
	}
}
