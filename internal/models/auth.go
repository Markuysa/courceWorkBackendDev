package models

type (
	GenerateOTPResponse struct {
		QR        []byte `json:"qr"`
		FailCause string `json:"fail_cause"`
	}
	GenerateOTPRequest struct {
		Username string `json:"username"`
	}
	ValidateOTPRequest struct {
		Username string `json:"username"`
		PassCode string `json:"passCode"`
	}
	ValidateOTPResponse struct {
		Passed bool `json:"passed"`
	}
	PrepareSignInRequest struct {
		Username string
		Password string
	}
	PrepareSignInResponse struct {
		AccessToken string
	}

	FinalizeSignInRequest struct {
		Username  string
		PassCode  string
		AccessKey string
	}
	FinalizeSignInResponse struct {
		SessionKey string
	}

	SaveSessionRequest struct {
		Username   string
		SessionKey string
	}
	SignUpRequest struct {
		Username string
		Password string
	}
	SignUpResponse struct {
		Success bool
	}

	SaveOTPRequest struct {
		Username string
		Secret   string
	}
	GetOTPRequest struct {
		Username string
	}
	GetUserRequest struct {
		Username string
	}

	SaveAccessKeyRequest struct {
		Key      string
		Username string
	}
	ValidateAccessKeyRequest struct {
		Key      string
		Username string
	}
	GetSessionRequest struct {
	}
)
