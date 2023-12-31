package users

import (
	"context"

	"github.com/Season-04/appos/core/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedUsersServiceServer
	repo *Repository
}

func NewServer(dataDir string) *Server {
	return &Server{
		repo: NewRepository(dataDir),
	}
}

func (s *Server) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &User{
		Name:         r.Name,
		EmailAddress: r.EmailAddress,
		Role:         RoleFromProtobuf[r.Role],
	}

	err := s.repo.CreateUser(user, r.Password)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: userToProtobuf(user),
	}, nil
}

func (s *Server) GetUserById(ctx context.Context, r *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user := s.repo.GetUserByID(r.Id)

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "There is no user with ID %v", r.Id)
	}

	return &pb.GetUserByIdResponse{
		User: userToProtobuf(user),
	}, nil
}

func (s *Server) GetUserByEmailAndPassword(ctx context.Context, r *pb.GetUserByEmailAndPasswordRequest) (*pb.GetUserByEmailAndPasswordResponse, error) {
	user := s.repo.GetUserByEmailAddressAndPassword(r.EmailAddress, r.Password)

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "There is no user with that email address and password")
	}

	return &pb.GetUserByEmailAndPasswordResponse{
		User: userToProtobuf(user),
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := s.repo.GetUserByID(r.Id)

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "There is no user with ID %v", r.Id)
	}

	if r.Name != nil {
		user.Name = r.Name.Value
	}

	if r.Role != pb.UserRole_USER_ROLE_UNSPECIFIED {
		user.Role = RoleFromProtobuf[r.Role]
	}

	err := s.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		User: userToProtobuf(user),
	}, nil
}

func (s *Server) ListUsers(ctx context.Context, r *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users := s.repo.ListAll()
	return &pb.ListUsersResponse{Users: usersToProtobuf(users)}, nil
}

var _ pb.UsersServiceServer = &Server{}

var RoleToProtobuf = map[UserRole]pb.UserRole{
	UserRoleAdmin: pb.UserRole_USER_ROLE_ADMIN,
	UserRoleUser:  pb.UserRole_USER_ROLE_USER,
}

var RoleFromProtobuf = map[pb.UserRole]UserRole{
	pb.UserRole_USER_ROLE_ADMIN: UserRoleAdmin,
	pb.UserRole_USER_ROLE_USER:  UserRoleUser,
}

func userToProtobuf(user *User) *pb.User {
	if user == nil {
		return nil
	}

	var lastSeenAt *timestamppb.Timestamp = nil
	if user.LastSeenAt != nil {
		lastSeenAt = timestamppb.New(*user.LastSeenAt)
	}

	return &pb.User{
		Id:           user.ID,
		Name:         user.Name,
		EmailAddress: user.EmailAddress,
		Role:         RoleToProtobuf[user.Role],
		LastSeenAt:   lastSeenAt,
	}
}

func usersToProtobuf(users []*User) []*pb.User {
	pbUsers := make([]*pb.User, len(users))

	for i, u := range users {
		pbUsers[i] = userToProtobuf(u)
	}

	return pbUsers
}
