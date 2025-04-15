package login

import (
	"context"
	"gin-jwt-auth/internal/login/dto"
)

type LoginService interface {
	LoginUser(ctx context.Context, req dto.LoginRequest) error
}

type loginService struct {
	repo LoginRepository
}

func NewLoginService(repo LoginRepository) LoginService {
	return &loginService{repo: repo}
}

func (s *loginService) LoginUser(ctx context.Context, req dto.LoginRequest) error {
	return s.repo.Login(ctx, req)
}
