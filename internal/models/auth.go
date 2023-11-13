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
	}
	ValidateOTPResponse struct {
	}
)
