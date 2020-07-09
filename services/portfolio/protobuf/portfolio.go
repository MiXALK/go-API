package portfolio

import (
    "context"
    "fmt"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

const (
    DB_PORTFOLIO_COLLECTION string = "portfolio"
    ACTION_STATUS_SUCCESS   string = "success"
)

type Server struct {
    Port     string
    DbClient *mongo.Client
}

type dataPortfolio struct {
    ID   primitive.ObjectID `bson:"_id,omitempty"`
    Name string             `bson:"name"`
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

    err = client.Connect(context.Background())
    if err != nil {
        return err
    }
    s.DbClient = client

    return nil
}

func (s *Server) DbDisconnect() {
    err := s.DbClient.Disconnect(context.Background())
    if err != nil {
        log.Printf("Error disconnect database: %v", err)
    }
}

func (s *Server) Create(_ context.Context, req *CreatePortfolioRequest) (*CreatePortfolioResponse, error) {
    res := &CreatePortfolioResponse{}
    if req.GetName() == "" {
        return res, status.Error(codes.InvalidArgument, "Name is empty")
    }

    portfolioCollection := s.getPortfolioCollection()
    portfolio := &Portfolio{
        Name: req.GetName(),
    }
    insertResult, err := portfolioCollection.InsertOne(context.Background(), portfolio)
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

    portfolioCollection := s.getPortfolioCollection()

    filter := bson.D{}
    var portfolios []*Portfolio

    cursor, err := portfolioCollection.Find(context.Background(), filter)
    if err != nil {
        return res, err
    }

    data := &dataPortfolio{}

    for cursor.Next(context.Background()) {
        err := cursor.Decode(data)

        portfolio := &Portfolio{
            Id: data.ID.Hex(),
            Name: data.Name,
        }

        if err != nil {
            return res, err
        }
        portfolios = append(portfolios, portfolio)
    }

    if err := cursor.Err(); err != nil {
        return res, err
    }

    if err := cursor.Close(context.Background()); err != nil {
        return res, err
    }

    res.Status = ACTION_STATUS_SUCCESS
    res.Portfolios = portfolios

    return res, nil
}

func (s *Server) Find(_ context.Context, req *FindPortfolioRequest) (*FindPortfolioResponse, error) {
    res := &FindPortfolioResponse{}

    id, err := primitive.ObjectIDFromHex(req.GetId())
    if err != nil {
        return res, err
    }

    portfolioCollection := s.getPortfolioCollection()
    data := &dataPortfolio{}
    filter := bson.M{"_id": id}
    err = portfolioCollection.FindOne(context.Background(), filter).Decode(data)
    if err != nil {
        return res, status.Error(codes.NotFound, fmt.Sprintf("Portfolio with id %s wasn't found", req.GetId()))
    }

    portfolio := &Portfolio{
        Id: data.ID.Hex(),
        Name: data.Name,
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

    portfolioCollection := s.getPortfolioCollection()
    id, err := primitive.ObjectIDFromHex(req.Id)
    if err != nil {
        return res, err
    }
    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{
        "name": req.Name,
    }}

    _, err = portfolioCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return res, err
    }

    res.Status = ACTION_STATUS_SUCCESS

    return res, nil
}

func (s *Server) Delete(_ context.Context, req *DeletePortfolioRequest) (*DeletePortfolioResponse, error) {
    res := &DeletePortfolioResponse{}

    id, err := primitive.ObjectIDFromHex(req.GetId())
    if err != nil {
        return res, err
    }

    portfolioCollection := s.getPortfolioCollection()
    filter := bson.M{"_id": id}
    _, err = portfolioCollection.DeleteOne(context.Background(), filter)
    if err != nil {
        return res, err
    }

    res.Status = ACTION_STATUS_SUCCESS

    return res, nil
}

func (s *Server) getPortfolioCollection() *mongo.Collection {
    return s.DbClient.Database(os.Getenv("DB_NAME")).Collection(DB_PORTFOLIO_COLLECTION)
}
