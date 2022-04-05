package usecase

import (
	"context"
	"fmt"

	"fairseller-backend/pkg/logger"
)

type AuthUseCase struct {
	userRepository userRepositoryInterface
	logger         logger.Interface
}

func NewAuthUseCase(userRepository userRepositoryInterface, l logger.Interface) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepository,
		logger:         l,
	}
}

// SignUpRequest - first step of sign up with sending code to phone.
func (uc *AuthUseCase) SignUpRequest(ctx context.Context, phone string) error {
	user, err := uc.userRepository.GetOneByPhone(phone)

	if err != nil {
		uc.logger.Error(err, "AuthUseCase - SignUpRequest")

		return err
	}

	if user.ID != 0 {
		return fmt.Errorf("user with phone %s is already exist", phone)
	}

	return nil
}
