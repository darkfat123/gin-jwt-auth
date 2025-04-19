package login

import (
	"context"
	"errors"
	"gin-jwt-auth/internal/login/dto"
	"gin-jwt-auth/model"
	"gin-jwt-auth/pkg/logger"
	"gin-jwt-auth/pkg/utils"

	"go.uber.org/zap"
)

type LoginService interface {
	LoginUser(ctx context.Context, req dto.LoginRequest) (*model.TokenPair, error)
}

type loginService struct {
	repo LoginRepository
}

func NewLoginService(repo LoginRepository) LoginService {
	return &loginService{repo: repo}
}

func (s *loginService) LoginUser(ctx context.Context, req dto.LoginRequest) (*model.TokenPair, error) {
	user, err := s.repo.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	isValid := utils.CheckPasswordHash(req.Password, user.Password)
	if !isValid {
		logger.Info("invalid password attempt", zap.String("username", req.Username))
		return nil, errors.New("invalid credentials")
	}

	tokens, err := utils.GenerateJWT(req.Username)
	if err != nil {
		return nil, err
	}

	logger.Info("login attempt for user", zap.String("username", req.Username))

	return tokens, nil
}
