package main

import (
	"fmt"
	portfolio "github.com/MiXALK/go-API/services/portfolio/protobuf"
	"log"
	"net"
	"os"
)

const SERVICE_NAME = "portfolio"

func main() {
	Port := os.Getenv("API_PORT")
	_, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &portfolio.Server{}

	err = s.DbConnect()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	//Test block
	//p := portfolio.CreatePortfolioRequest{
	//	Name: "TEST",
	//}
	//_, er := s.Create(context.TODO(), &p)
	//if er != nil {
	//	log.Fatalf("failed: %v", err)
	//}
	//
	//a := portfolio.GetAllPortfolioRequest{}
	//_, err = s.GetAll(context.TODO(), &a)
	//if err != nil {
	//	log.Fatalf("failed: %v", err)
	//}
	//a := portfolio.UpdatePortfolioRequest{
	//	Id: "5ec67b0509a4edc79d1f8312",
	//	Name: "Changed Name",
	//}
	//res, err := s.Update(context.TODO(), &a)
	//if err != nil {
	//	log.Fatalf("failed: %v", err)
	//}
	//fmt.Println(res)

	defer s.DbDisconnect()

	log.Printf(
		"%s service started on  %s:%s (localhost:%s)",
		SERVICE_NAME,
		os.Getenv("PORTFOLIO_HOST"),
		Port,
		os.Getenv("PORTFOLIO_PORT_LOCAL"),
	)
}
