package service

import (
	"context"
	"github.com/google/uuid"
	"petProject/model"
)

// UserService is a service for users
//go:generate mockery --dir . --name UserService --output ./mocks
type UserService interface {
	GetUser(ctx context.Context, uuid uuid.UUID) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, uuid uuid.UUID) error
}

// For example
// FileMetaService is a service for files
//g!o:generate mockery --dir . --name FileMetaService --output ./mocks
//type FileMetaService interface {
//	GetFileMeta(context.Context, uuid.UUID) (*model.File, error)
//	CreateFileMeta(context.Context, *model.File) (*model.File, error)
//	UpdateFileMeta(context.Context, *model.File) (*model.File, error)
//	DeleteFileMeta(context.Context, uuid.UUID) error
//}