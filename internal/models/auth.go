package models

type (
	GenerateOTPResponse struct {
		QR        []byte `json:"qr"`
		Secret    string `json:"secret"`
		FailCause string `json:"fail_cause"`
	}
	GenerateOTPRequest struct {
		Username string `validate:"required" json:"username"`
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
		FailCause   string
	}

	FinalizeSignInRequest struct {
		Username  string `json:"username"`
		PassCode  string `json:"pass_code"`
		AccessKey string `json:"access_key"`
	}
	FinalizeSignInResponse struct {
		SessionKey string `json:"session_key"`
		FailCause  string `json:"fail_cause"`
	}

	SaveSessionRequest struct {
		Username   string
		SessionKey string
		Role       int
	}
	SignUpRequest struct {
		Username string
		Password string
	}
	SignUpResponse struct {
		Success   bool   `json:"success"`
		QR        []byte `json:"qr"`
		FailCause string `json:"fail_cause"`
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
		SessionKey string
	}

	Session struct {
		UserID int
		Role   int
	}
)
