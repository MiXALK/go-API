package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"

    portfolio "github.com/MiXALK/go-API/services/api-client/services/portfolio/protobuf"
    user "github.com/MiXALK/go-API/services/api-client/services/user/protobuf"
)

var (
    portfolioServerAddress = fmt.Sprintf(
        "%s:%s",
        os.Getenv("PORTFOLIO_HOST"),
        os.Getenv("PORTFOLIO_PORT"),
    )
    userServerAddress = fmt.Sprintf(
        "%s:%s",
        os.Getenv("USER_HOST"),
        os.Getenv("USER_PORT"),
    )
)

func main() {
	proxyAddr := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	HTTPProxy(proxyAddr)
}

func HTTPProxy(proxyAddr string) {
    grpcGwMux := runtime.NewServeMux()

    grpcPortfolioConn, err := grpc.Dial(
        portfolioServerAddress,
        grpc.WithInsecure(),
    )
    if err != nil {
        log.Fatalln("Filed to connect to Portfolio service", err)
    }
    defer grpcPortfolioConn.Close()

    err = portfolio.RegisterPortfolioServiceHandler(
        context.Background(),
        grpcGwMux,
        grpcPortfolioConn,
    )
    if err != nil {
        log.Fatalln("Filed to start portfolio HTTP server", err)
    }

    grpcUserConn, err := grpc.Dial(
        userServerAddress,
        grpc.WithInsecure(),
    )
    if err != nil {
        log.Fatalln("Filed to connect to User service", err)
    }
    defer grpcUserConn.Close()

    err = user.RegisterUserServiceHandler(
        context.Background(),
        grpcGwMux,
        grpcUserConn,
    )

    if err != nil {
        log.Fatalln("Filed to start user HTTP server", err)
    }

    mux := http.NewServeMux()
    mux.Handle("/v1/", grpcGwMux)

    log.Printf(
        "Starting HTTP server at %s",
        proxyAddr,
    )
    log.Fatal(http.ListenAndServe(proxyAddr, mux))
}
