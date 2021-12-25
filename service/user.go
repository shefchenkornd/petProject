package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"petProject/lib/types"
	"petProject/model"
	"petProject/store"
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

//GetUser get user by ID
func (s UserWebService) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	userDB, err := s.store.User.GetUser(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetUser")
	}
	if userDB == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("User '%s' not found", userID.String()))
	}

	return userDB.ToWeb(), nil
}

// CreateUser creates user
func (s UserWebService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.ID = uuid.New()

	createdDBUser, err := s.store.User.CreateUser(ctx, user.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.CreateUser error")
	}

	return createdDBUser.ToWeb(), nil
}

// UpdateUser creates user
func (s UserWebService) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if user == nil {
		return nil, errors.New("svc.user.UpdateUser error")
	}

	updatedUser, err := s.store.User.UpdateUser(ctx, user.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.UpdateUser error")
	}

	return updatedUser.ToWeb(), nil
}

func (s UserWebService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	// Check if user exists
	userDB, err := s.store.User.GetUser(ctx, userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.GetUser error")
	}
	if userDB == nil {
		return errors.Wrap(types.ErrNotFound, fmt.Sprintf("User '%s' not found", userID.String()))
	}

	err = s.store.User.DeleteUser(ctx, userDB.ID)
	if err != nil {
		return errors.Wrap(err, "svc.user.DeleteUser error")
	}

	return nil
}
