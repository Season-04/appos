package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/staugaard/app-os/core/internal/graph/model"
	"github.com/staugaard/app-os/core/internal/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.User, error) {
	if currentUser := r.CurrentUser(ctx); currentUser == nil || currentUser.Role != pb.UserRole_USER_ROLE_ADMIN {
		return nil, errAccessDenied
	}

	request := &pb.CreateUserRequest{
		Name:         input.Name,
		EmailAddress: input.EmailAddress,
		Password:     input.Password,
		Role:         model.RoleToProtobuf[input.Role],
	}

	response, err := r.UsersService.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return &model.User{User: response.User}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUserInput) (*model.User, error) {
	currentUser := r.CurrentUser(ctx)
	if currentUser == nil {
		return nil, errors.Wrap(errAccessDenied, "Not Authenticated")
	}

	intId, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		return nil, nil
	}

	if currentUser.Role != pb.UserRole_USER_ROLE_ADMIN && currentUser.Id != uint32(intId) {
		return nil, errors.Wrap(errAccessDenied, "Only admins can edit other users")
	}

	request := &pb.UpdateUserRequest{
		Id: uint32(intId),
	}

	if input.Name != nil {
		request.Name = wrapperspb.String(*input.Name)
	}

	if input.Role != nil {
		if currentUser.Role != pb.UserRole_USER_ROLE_ADMIN {
			return nil, errors.Wrap(errAccessDenied, "You can not edit your wn role")
		}
		request.Role = model.RoleToProtobuf[*input.Role]
	}

	response, err := r.UsersService.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return &model.User{User: response.User}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	if r.CurrentUserID(ctx) == 0 {
		return nil, errAccessDenied
	}

	log.Println("Graph got user", r.CurrentUser(ctx))
	response, err := r.UsersService.List(ctx, &pb.ListRequest{})
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, len(response.Users))
	for i, user := range response.Users {
		users[i] = &model.User{User: user}
	}
	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	if r.CurrentUserID(ctx) == 0 {
		return nil, errAccessDenied
	}

	intId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, nil
	}

	response, err := r.UsersService.GetById(ctx, &pb.GetUserByIdRequest{Id: uint32(intId)})
	if err != nil {
		return nil, err
	}

	if response.User == nil {
		return nil, nil
	}

	return &model.User{User: response.User}, nil
}

// Role is the resolver for the role field.
func (r *userResolver) Role(ctx context.Context, obj *model.User) (model.UserRole, error) {
	switch obj.Role {
	case pb.UserRole_USER_ROLE_ADMIN:
		return model.UserRoleAdmin, nil
	case pb.UserRole_USER_ROLE_USER:
		return model.UserRoleUser, nil
	default:
		return model.UserRoleUser, errors.New("unsupported user role")
	}
}

// LastSeenAt is the resolver for the lastSeenAt field.
func (r *userResolver) LastSeenAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	if obj.LastSeenAt == nil {
		return nil, nil
	}

	lastSeenAt := obj.LastSeenAt.AsTime()
	return &lastSeenAt, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
