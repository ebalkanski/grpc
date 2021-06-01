package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/ebalkanski/grpc/proto"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUsersClient(conn)

	req := &pb.CreateUserRequest{User: &pb.User{
		Name: "John",
	}}
	_, err = client.CreateUser(context.Background(), req)
	if err != nil {
		fmt.Printf("error creating user: %v\n", err)
	}

	gResp, err := client.GetUsers(context.Background(), &pb.GetUsersRequest{})
	if err != nil {
		fmt.Printf("error getting users: %v\n", err)
	}

	for _, user := range gResp.GetUsers() {
		fmt.Printf("user = %#v\n", user)
	}
}
