package register

import (
	"context"
	"errors"
	"gin-jwt-auth/internal/register/dto"
	"gin-jwt-auth/pkg/logger"
	"gin-jwt-auth/pkg/utils"

	"go.uber.org/zap"
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
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		return errors.New("unable to process password")
	}

	duplicateCheckErr := s.repo.CheckDuplicate(ctx, req)
	if duplicateCheckErr != nil {
		return errors.New("username or email already exists")
	}

	hashedReq := dto.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}

	logger.Info("user registered successfully", zap.String("username", req.Username))

	return s.repo.Register(ctx, hashedReq)
}
