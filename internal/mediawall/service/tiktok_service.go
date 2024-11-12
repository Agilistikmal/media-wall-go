package service

import (
	"net/url"

	"github.com/agilistikmal/media-wall-go/internal/pkg"
	"github.com/spf13/viper"
)

type TikTokService struct {
}

func NewTikTokService() *TikTokService {
	return &TikTokService{}
}

func (s *TikTokService) GetOAuthURL() string {
	url, _ := url.Parse("https://www.tiktok.com/v2/auth/authorize/")

	query := url.Query()

	query.Add("client_key", viper.GetString("tiktok.client_key"))
	query.Add("response_type", "code")
	query.Add("scope", "user.info.basic")
	query.Add("redirect_uri", viper.GetString("tiktok.redirect_url"))
	query.Add("state", pkg.RandString(4))

	url.RawQuery = query.Encode()

	return url.String()
}
