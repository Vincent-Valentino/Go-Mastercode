package repository

import (
	"context"
	"errors"
	"time"
	"internal/model/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
}

type UserRepository struct {
	User []User
}