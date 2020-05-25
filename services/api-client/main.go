package main

import (
	"context"
	"fmt"
	portfolioService "github.com/MiXALK/go-API/services/portfolio/protobuf"
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

	portfolioClient := portfolioService.NewPortfolioServiceClient(grpcPortfolioConn)

	mux := http.NewServeMux()

	p := &portfolioService.CreatePortfolioRequest{
		Name: "New",
	}

	fmt.Println(portfolioClient.Create(context.TODO(), p))

	log.Printf(
		"tarting HTTP server at %s",
		proxyAddr,
	)
	log.Fatal(http.ListenAndServe(proxyAddr, mux))
}
