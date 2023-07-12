package register

type CreateEmailVerification struct {
	SUBJECT           string
	EMAIL             string
	VERIFICATION_CODE string
}

type VerifyEmail struct {
	EMAIL             string `json:"email" binding:"required"`
	VERIFICATION_CODE string `json:"verification_code" binding:"required"`
}
