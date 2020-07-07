package main

import (
	"fmt"
	user "github.com/MiXALK/go-API/services/user/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const SERVICE_NAME string = "user_server"

func main() {
	Port := os.Getenv("USER_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &user.Server{}

	err = s.DbConnect()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer s.DbDisconnect()

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, s)

	log.Printf(
		"%s service started on  %s:%s (localhost:%s)",
		SERVICE_NAME,
		os.Getenv("USER_HOST"),
		Port,
		os.Getenv("USER_PORT"),
	)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
