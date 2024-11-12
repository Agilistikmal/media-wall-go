package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/agilistikmal/media-wall-go/internal/mediawall/model"
	"github.com/agilistikmal/media-wall-go/internal/pkg"
	"github.com/spf13/viper"
)

type TikTokService struct {
}

func NewTikTokService() *TikTokService {
	return &TikTokService{}
}

func (s *TikTokService) GetOAuthURL() string {
	baseUrl, _ := url.Parse("https://www.tiktok.com/v2/auth/authorize/")

	query := baseUrl.Query()

	query.Add("client_key", viper.GetString("tiktok.client_key"))
	query.Add("response_type", "code")
	query.Add("scope", "user.info.basic,video.list")
	query.Add("redirect_uri", viper.GetString("tiktok.redirect_url"))
	query.Add("state", pkg.RandString(4))

	baseUrl.RawQuery = query.Encode()

	return baseUrl.String()
}

func (s *TikTokService) GetAccessToken(code string) (*model.OAuthResponse, error) {
	baseUrl, _ := url.Parse("https://open.tiktokapis.com/v2/oauth/token/")

	data := url.Values{}

	data.Set("client_key", viper.GetString("tiktok.client_key"))
	data.Set("client_secret", viper.GetString("tiktok.client_secret"))
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", viper.GetString("tiktok.redirect_url"))

	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, baseUrl.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var oAuthResponseError *model.OAuthResponseError
	json.Unmarshal(body, &oAuthResponseError)

	if oAuthResponseError.Error != "" {
		return nil, fmt.Errorf("error: %s - %s", oAuthResponseError.Error, oAuthResponseError.ErrorDescription)
	}

	var oAuthResponse *model.OAuthResponse
	json.Unmarshal(body, &oAuthResponse)

	return oAuthResponse, nil
}

func (s *TikTokService) QueryVideos(accessToken string) (*model.QueryUserVideoResponse, error) {
	baseUrl, _ := url.Parse("https://open.tiktokapis.com/v2/video/query/")

	query := baseUrl.Query()

	query.Add("fields", "id,title")

	baseUrl.RawQuery = query.Encode()

	requestBody := []byte(`{ 
		"filters": { 
			"video_ids": [
				"1234123412345678567",
				"1010102020203030303"
			]
		}
	}`)

	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, baseUrl.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var queryUserVideoResponse *model.QueryUserVideoResponse
	json.Unmarshal(body, &queryUserVideoResponse)

	return queryUserVideoResponse, nil
}
