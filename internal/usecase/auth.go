package usecase

import (
	"fmt"

	"fairseller-backend/internal/entity"
)

type AuthUseCase struct {
	userRepository userRepositoryInterface
	logger         entity.Logger
}

func NewAuthUseCase(userRepository userRepositoryInterface, l entity.Logger) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepository,
		logger:         l,
	}
}

func (uc *AuthUseCase) SignUpRequest(phone string) error {
	user, err := uc.userRepository.GetOneByPhone(phone)

	if err != nil {
		uc.logger.DatabaseError(err)

		return err
	}

	if user.ID != 0 {
		err = fmt.Errorf("user with phone %s is already exist", phone)
		uc.logger.UseCaseError(err)

		return err
	}

	return nil
}
