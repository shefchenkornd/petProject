package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"petProject/model"
	"github.com/google/uuid"
)

// UserWebService ...
type UserWebService struct {
	ctx   context.Context
	store *store.Store
}

// NewUserWebService creates a new user web service
func NewUserWebService(ctx context.Context, store *store.Store) *UserWebService {
	return &UserWebService{
		ctx:   ctx,
		store: store,
	}
}


func (s UserWebService) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	userDB, err := s.store.GetUser(ctx, userID)
	if err != nil {
		return nil,  errors.Wrap(err, "svc.user.GetUser")
	}
	if userDB == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("User '%s' not found", userID.String()))
	}

	return userDB.ToWeb(), nil
}
