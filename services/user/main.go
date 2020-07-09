package main

import (
    "fmt"
    "log"
    "net"
    "os"

    "google.golang.org/grpc"

    "github.com/MiXALK/go-API/services/user/protobuf"
)

func main() {
    port := os.Getenv("USER_PORT")
    listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
    if err != nil {
        log.Fatalf("Failed to listen on port %s: %v", port, err)
    }

    s := &user.Server{}

    err = s.DbConnect()
    if err != nil {
        log.Fatalf("Failed to connect database: %v", err)
    }
    defer s.DbDisconnect()

    grpcServer := grpc.NewServer()
    user.RegisterUserServiceServer(grpcServer, s)
    err = grpcServer.Serve(listener)
    if err != nil {
        log.Fatalf("failed to serve: %s", err)
    }

    log.Printf(
        "%s service succesfully started on  %s:%s (localhost:%s)",
        os.Getenv("USER_HOST"),
        os.Getenv("USER_HOST"),
        port,
        port,
    )
}
