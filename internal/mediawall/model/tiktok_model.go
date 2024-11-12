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

type QueryUserVideoResponse struct {
	Data  string              `json:"data,omitempty"`
	Error TikTokErrorResponse `json:"error,omitempty"`
}

type QueryUserVideoData struct {
	Videos []VideoData `json:"videos,omitempty"`
}

type VideoData struct {
	Title string `json:"title,omitempty"`
	ID    string `json:"id,omitempty"`
}

type TikTokErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	LogID   string `json:"log_id,omitempty"`
}
