package service

import "net/url"

type TikTokService struct {
}

func NewTikTokService() *TikTokService {
	return &TikTokService{}
}

func (s *TikTokService) Authorize() {
	url, _ := url.Parse("https://www.tiktok.com/v2/auth/authorize/")

	url.Query().Add("client_key", "")
	url.Query().Add("response_type", "")
	url.Query().Add("scope", "")
	url.Query().Add("redirect_uri", "")
	url.Query().Add("state", "")
}
