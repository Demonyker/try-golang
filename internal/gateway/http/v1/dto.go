package v1

type signUpRequestBody struct {
	Phone string `json:"phone" binding:"required,e164"`
}
