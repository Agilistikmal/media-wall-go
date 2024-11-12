package model

type OAuthResponse struct {
	OpenID           string `json:"open_id,omitempty"`
	Scope            string `json:"scope,omitempty"`
	AccessToken      string `json:"access_token,omitempty"`
	ExpiresIn        int64  `json:"expires_in,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	RefreshExpiresIn int64  `json:"refresh_expires_in,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
}

type OAuthResponseError struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	LogID            string `json:"log_id,omitempty"`
}
