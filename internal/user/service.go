package user

import (
	"context"
	"gin-jwt-auth/internal/user/dto"
)

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*dto.UserResponse, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*dto.UserResponse, error) {
	return s.repo.GetUserByID(ctx, id)
}
