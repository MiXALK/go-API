package main

import (
	"fmt"
	portfolio "github.com/MiXALK/go-API/services/portfolio/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const SERVICE_NAME = "portfolio_server"

func main() {
	Port := os.Getenv("PORTFOLIO_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &portfolio.Server{}

	err = s.DbConnect()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer s.DbDisconnect()

	grpcServer := grpc.NewServer()
	portfolio.RegisterPortfolioServiceServer(grpcServer, s)

	log.Printf(
		"%s service started on  %s:%s (localhost:%s)",
		SERVICE_NAME,
		os.Getenv("PORTFOLIO_HOST"),
		Port,
		os.Getenv("PORTFOLIO_PORT_LOCAL"),
	)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
