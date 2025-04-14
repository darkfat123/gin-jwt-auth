package service

import (
	"context"
	"gin-jwt-auth/internal/model"
	"gin-jwt-auth/internal/repository"
)

type RegisterService interface {
	RegisterUser(ctx context.Context, req model.RegisterRequest) error
}

type registerService struct {
	repo repository.RegisterRepository
}

func NewRegisterService(repo repository.RegisterRepository) RegisterService {
	return &registerService{repo: repo}
}

func (s *registerService) RegisterUser(ctx context.Context, req model.RegisterRequest) error {
	return s.repo.Register(ctx, req)
}
