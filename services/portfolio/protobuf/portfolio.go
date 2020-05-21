package portfolio

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

const (
	DB_PORTFOLIO_COLLECTION string = "portfolio"
	ACTION_STATUS_SUCCESS   string = "success"
)

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

	var portfolioId string
	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if ok {
		portfolioId = oid.Hex()
	} else {
		return res, status.Error(codes.Internal, "Error portfolio saving to database")
	}

	res.Status = ACTION_STATUS_SUCCESS
	res.ID = portfolioId

	return res, nil
}

func (s *Server) GetAll(_ context.Context, _ *GetAllPortfolioRequest) (*GetAllPortfolioResponse, error) {
	res := &GetAllPortfolioResponse{}

	collection := s.DbClient.Database(os.Getenv("DB_NAME")).Collection(DB_PORTFOLIO_COLLECTION)

	filter := bson.D{}
	var results []*Portfolio

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return res, err
	}

	for cur.Next(context.TODO()) {
		portfolio := &Portfolio{}
		err := cur.Decode(portfolio)
		if err != nil {
			return res, err
		}
		results = append(results, portfolio)
	}

	if err := cur.Err(); err != nil {
		return res, err
	}

	if err := cur.Close(context.TODO()); err != nil {
		return res, err
	}

	res.Status = ACTION_STATUS_SUCCESS
	res.Portfolios = results

	return res, nil
}

func (s *Server) Find(_ context.Context, req *FindPortfolioRequest) (*FindPortfolioResponse, error) {
	res := &FindPortfolioResponse{}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return res, err
	}

	collection := s.DbClient.Database(os.Getenv("DB_NAME")).Collection(DB_PORTFOLIO_COLLECTION)
	portfolio := &Portfolio{}
	filter := bson.M{"_id": id}
	err = collection.FindOne(context.TODO(), filter).Decode(portfolio)
	if err != nil {
		return res, status.Error(codes.NotFound, fmt.Sprintf("Portfolio with id %s wasn't found", req.ID))
	}

	res.Status = ACTION_STATUS_SUCCESS
	res.Portfolio = portfolio

	return res, nil
}

func (s *Server) Update(_ context.Context, req *UpdatePortfolioRequest) (*UpdatePortfolioResponse, error) {
	res := &UpdatePortfolioResponse{}

	if req.Name == "" {
		return res, status.Error(codes.InvalidArgument, "Name is empty")
	}

	collection := s.DbClient.Database(os.Getenv("DB_NAME")).Collection(DB_PORTFOLIO_COLLECTION)
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return res, err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"name": req.Name,
	}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return res, err
	}

	res.Status = ACTION_STATUS_SUCCESS

	return res, nil
}
