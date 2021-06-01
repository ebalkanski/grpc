package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ebalkanski/grpc/internal/service"
	pb "github.com/ebalkanski/grpc/proto"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	svc := service.New()

	s := grpc.NewServer()
	pb.RegisterUsersServer(s, svc)

	log.Println("starting users grpc server...")
	if err := s.Serve(ln); err != nil {
		log.Fatal(err)
	}

	log.Println("bye bye")
}
