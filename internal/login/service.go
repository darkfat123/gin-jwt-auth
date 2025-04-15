package login

import (
	"context"
	"gin-jwt-auth/internal/login/dto"
	"gin-jwt-auth/pkg/utils"
)

type LoginService interface {
	LoginUser(ctx context.Context, req dto.LoginRequest) (string, error)
}

type loginService struct {
	repo LoginRepository
}

func NewLoginService(repo LoginRepository) LoginService {
	return &loginService{repo: repo}
}

func (s *loginService) LoginUser(ctx context.Context, req dto.LoginRequest) (string, error) {
	err := s.repo.Login(ctx, req)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(req.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
