package portfolio

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

const DB_PORTFOLIO_COLLECTION string = "portfolio"

type Server struct {
	Port     string
	DbClient *mongo.Client
}

func (s *Server) DbConnect() error {
	var client *mongo.Client

	strURI := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(strURI))
	if err != nil {
		return err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}
	s.DbClient = client
	return nil
}

func (s *Server) DbDisconnect() {
	err := s.DbClient.Disconnect(context.TODO())
	if err != nil {
		log.Printf("Error disconnect database: %v", err)
	}
}

func (s *Server) Create(_ context.Context, req *CreatePortfolioRequest) (*CreatePortfolioResponse, error) {
	res := &CreatePortfolioResponse{}
	if req.Name == "" {
		return res, status.Error(codes.InvalidArgument, "Name is empty")
	}

	collection := s.DbClient.Database(os.Getenv("DB_NAME")).Collection(DB_PORTFOLIO_COLLECTION)
	portfolio := &Portfolio{
		Name: req.Name,
	}
	insertResult, err := collection.InsertOne(context.TODO(), portfolio)
	if err != nil {
		return res, err
	}

	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if ok {
		portfolio.ID = oid.Hex()
	} else {
		return res, status.Error(codes.Internal, "Error portfolio saving to database")
	}

	res.Status = "success"
	res.ID = portfolio.GetID()

	//fmt.Println(res)

	return res, nil
}
