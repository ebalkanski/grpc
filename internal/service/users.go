package service

import (
	"context"
	"fmt"

	pb "github.com/ebalkanski/grpc/proto"
)

type Users struct {
	pb.UnimplementedUsersServer
	Users []*pb.User
}

func New() *Users {
	return &Users{
		Users: []*pb.User{
			{
				Name: "Bob",
			},
		},
	}
}

func (svc *Users) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Users: svc.Users,
	}, nil
}

func (svc *Users) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	for _, u := range svc.Users {
		if request.GetUser().GetName() == u.GetName() {
			return nil, fmt.Errorf("user %q already exists", u.GetName())
		}
	}
	svc.Users = append(svc.Users, request.GetUser())

	return &pb.CreateUserResponse{}, nil
}
