package user

import (
    "context"
    "fmt"
    "log"
    "os"
    "regexp"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "github.com/MiXALK/go-API/services/user/app"
)

const (
    DB_USER_COLLECTION string = "user"
    ACTION_STATUS_SUCCESS   string = "success"
)

type Server struct {
    Port     string
    DbClient *mongo.Client
}

type dataUser struct {
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

func (s *Server) Create(_ context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
    res := &CreateUserResponse{}
    err := req.validateUserRequest()
    if err != nil {
        return res, err
    }

    userCollection := s.getUserCollection()
    user := &User{
        Name: req.GetName(),
        Email: req.GetEmail(),
        Password: req.GetPassword(),
    }
    insertResult, err := userCollection.InsertOne(context.Background(), user)
    if err != nil {
        return res, err
    }

    var userId string
    oid, ok := insertResult.InsertedID.(primitive.ObjectID)
    if ok {
        userId = oid.Hex()
    } else {
        return res, status.Error(codes.Internal, "Error user saving to database")
    }

    res.Status = ACTION_STATUS_SUCCESS
    res.Id = userId

    return res, nil
}

func (s *Server) getUserCollection() *mongo.Collection {
    return s.DbClient.Database(os.Getenv("DB_NAME")).Collection(DB_USER_COLLECTION)
}

//TODO Use go-playground/validator for validation
var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (req *CreateUserRequest) validateUserRequest() error {
    if req.GetName() == "" {
        return app.ErrNameIsEmpty
    }

    if req.GetPassword() == "" {
        return app.ErrPasswordIsEmpty
    }

    if !rxEmail.MatchString(req.GetEmail()) {
        return app.ErrEmailIsEmpty
    }

    return nil
}
