package app

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

var (
    ErrEmailIsEmpty = status.Error(codes.InvalidArgument, "Not a valid email address")
    ErrPasswordIsEmpty = status.Error(codes.InvalidArgument, "Password is empty")
    ErrNameIsEmpty = status.Error(codes.InvalidArgument, "Name is empty")
)
