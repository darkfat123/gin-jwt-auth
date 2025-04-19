package refresh

import (
	"context"
	"errors"
	"gin-jwt-auth/model"
	"gin-jwt-auth/pkg/logger"
	"gin-jwt-auth/pkg/utils"
	"log"

	"go.uber.org/zap"
)

type RefreshService interface {
	RefreshUser(ctx context.Context, reqToken string) (*model.TokenPair, error)
}

type refreshService struct {
	repo RefreshRepository
}

func NewRefreshService(repo RefreshRepository) RefreshService {
	return &refreshService{repo: repo}
}
func (s *refreshService) RefreshUser(ctx context.Context, reqToken string) (*model.TokenPair, error) {
	claims, err := utils.DecodeJWT(reqToken)
	if err != nil {
		log.Println("Error:", err)
		return nil, errors.New("unable to decode jwt")
	}

	username, ok := claims["iss"].(string)
	if !ok {
		log.Println("username is not a string")
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		log.Println("jti is not a string")
	}

	id, err := s.repo.Refresh(ctx, jti, username)
	if err != nil {
		log.Println("failed to hash password", err)
		return nil, errors.New("unable to process password")
	}

	newClaims, err := utils.GenerateJWT(username)
	if err != nil {
		logger.Error("failed to generate jwt", zap.Error(err))
		return nil, errors.New("unable to process jwt")
	}

	err = s.repo.UpdateRefreshToken(ctx, newClaims.RefreshToken, id)
	if err != nil {
		log.Println("failed to update refresh token", err)
		return nil, err
	}

	logger.Info("user refreshed successfully")

	return newClaims, nil
}
