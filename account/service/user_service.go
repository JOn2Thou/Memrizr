package service

import (
	"context"
	"github.com/JOn2Thou/memrizr/account/model"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}

func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	panic("implement me")
}
