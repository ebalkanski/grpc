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

	names := []string{"John", "Bob", "Bill", "John"}
	for _, name := range names {
		req := &pb.CreateUserRequest{
			Name: name,
		}
		_, err = client.CreateUser(context.Background(), req)
		if err != nil {
			fmt.Printf("error creating user: %v\n", err)
		}
	}

	gResp, err := client.GetUsers(context.Background(), &pb.GetUsersRequest{})
	if err != nil {
		fmt.Printf("error getting users: %v\n", err)
	}

	for _, user := range gResp.GetUsers() {
		fmt.Printf("user = %#v\n", user)
	}
}
