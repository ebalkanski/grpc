package service

import (
	"context"
	"fmt"

	pb "github.com/ebalkanski/grpc/proto"
)

type Users struct {
	pb.UnimplementedUsersServer
	Users      []*pb.User
	LastUserId int
}

func New() *Users {
	return &Users{
		Users: []*pb.User{
			{
				Id:   1,
				Name: "Default",
			},
		},
		LastUserId: 1,
	}
}

func (svc *Users) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Users: svc.Users,
	}, nil
}

func (svc *Users) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	for _, u := range svc.Users {
		if request.GetName() == u.GetName() {
			return nil, fmt.Errorf("user %q already exists", request.GetName())
		}
	}

	svc.LastUserId += 1
	user := &pb.User{
		Id:   int32(svc.LastUserId),
		Name: request.Name,
	}
	svc.Users = append(svc.Users, user)

	return &pb.CreateUserResponse{
		User: user,
	}, nil
}
