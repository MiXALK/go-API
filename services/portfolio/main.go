package main

import (
    "fmt"
    "log"
    "net"
    "os"

    "google.golang.org/grpc"

    "github.com/MiXALK/go-API/services/portfolio/protobuf"
)

func main() {
    port := os.Getenv("PORTFOLIO_PORT")
    listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
    if err != nil {
        log.Fatalf("Failed to listen on port %s: %v", port, err)
    }

    s := &portfolio.Server{}

    err = s.DbConnect()
    if err != nil {
        log.Fatalf("Failed to connect database: %v", err)
    }
    defer s.DbDisconnect()

    grpcServer := grpc.NewServer()
    portfolio.RegisterPortfolioServiceServer(grpcServer, s)
    err = grpcServer.Serve(listener)
    if err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }

    log.Printf(
        "%s service succesfully started on %s:%s (localhost:%s)",
        os.Getenv("PORTFOLIO_HOST"),
        os.Getenv("PORTFOLIO_HOST"),
        port,
        port,
    )
}
