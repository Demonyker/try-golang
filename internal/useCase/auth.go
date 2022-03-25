package useCase

import (
	"context"
	"fairseller-backend/internal/entity"
	"fairseller-backend/pkg/logger"
	"fmt"
)

// AuthUseCase -.
type AuthUseCase struct {
	userRepository UserRepository
	logger         logger.Interface
}

func NewAuthUseCase(userRepository UserRepository, logger logger.Interface) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepository,
		logger:         logger,
	}
}

// SignUpRequest - first step of sign up with sending code to phone
func (uc *AuthUseCase) SignUpRequest(ctx context.Context, dto entity.SignUpRequest) error {
	user, err := uc.userRepository.GetOneByPhone(ctx, dto.Phone)
	if err != nil {
		uc.logger.Error(err, "AuthUseCase - SignUpRequest")
		return fmt.Errorf("internal server error")
	}

	if user.ID != 0 {
		return fmt.Errorf("user with phone %s is already exist", dto.Phone)
	}

	return nil
}
