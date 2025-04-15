package register

import (
	"context"
	"gin-jwt-auth/internal/register/dto"
)

type RegisterService interface {
	RegisterUser(ctx context.Context, req dto.RegisterRequest) error
}

type registerService struct {
	repo RegisterRepository
}

func NewRegisterService(repo RegisterRepository) RegisterService {
	return &registerService{repo: repo}
}

func (s *registerService) RegisterUser(ctx context.Context, req dto.RegisterRequest) error {
	return s.repo.Register(ctx, req)
}
