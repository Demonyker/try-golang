package v1

type SignUpRequestDto struct {
	Phone string `json:"phone" binding:"required,e164"`
}
