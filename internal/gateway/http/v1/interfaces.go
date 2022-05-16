package v1

type (
	AuthUseCase interface {
		SignUpRequest(phone string) error
	}
)
