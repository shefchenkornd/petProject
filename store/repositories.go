package store

import (
	"context"
	"github.com/google/uuid"
	"petProject/model"
)

// UserRepo is a store for users
//go:generate mockery --dir . --name UserRepo --output ./mocks
type UserRepo interface {
	GetUser(ctx context.Context, uuid uuid.UUID) (*model.DBUser, error)
	CreateUser(ctx context.Context, user *model.DBUser) (*model.DBUser, error)
	UpdateUser(ctx context.Context, user *model.DBUser) (*model.DBUser, error)
	DeleteUser(ctx context.Context, uuid uuid.UUID) error
}