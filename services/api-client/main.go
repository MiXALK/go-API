package main

import (
	"context"
	"fmt"
	portfolio "github.com/MiXALK/go-API/services/api-client/services/portfolio/protobuf"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

var (
	portfolioServerAddress = fmt.Sprintf(
		"%s:%s",
		os.Getenv("PORTFOLIO_HOST"),
		os.Getenv("PORTFOLIO_PORT"),
	)
)

func main() {
	proxyAddr := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	HTTPProxy(proxyAddr)
}

func HTTPProxy(proxyAddr string) {
	grpcPortfolioConn, err := grpc.Dial(
		portfolioServerAddress,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Filed to connect to Portfolio service", err)
	}
	defer grpcPortfolioConn.Close()

	grpcGwMux := runtime.NewServeMux()

	err = portfolio.RegisterPortfolioServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcPortfolioConn,
	)

	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/", grpcGwMux)

	log.Printf(
		"Starting HTTP server at %s",
		proxyAddr,
	)
	log.Fatal(http.ListenAndServe(proxyAddr, mux))
}
