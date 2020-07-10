package app

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

var (
    ErrInvalidEmail    = status.Error(codes.InvalidArgument, "Not a valid email address")
    ErrInvalidPhone    = status.Error(codes.InvalidArgument, "Not a valid phone")
    ErrPasswordIsEmpty = status.Error(codes.InvalidArgument, "Password is empty")
    ErrNameIsEmpty = status.Error(codes.InvalidArgument, "Name is empty")
)
