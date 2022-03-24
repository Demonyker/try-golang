package useCase

import (
	"context"
	"fairseller-backend/internal/entity"
)

// AuthUseCase -.
type AuthUseCase struct {
	userRepository UserRepository
}

func NewAuthUseCase(userRepository UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepository,
	}
}

// SignUpRequest - first step of sign up with sending code to phone
func (uc *AuthUseCase) SignUpRequest(ctx context.Context, dto entity.SignUpRequestDto) error {
	//if err != nil {
	//	return nil, fmt.Errorf("AuthUseCase - History - s.repo.GetHistory: %w", err)
	//}

	return nil
}
