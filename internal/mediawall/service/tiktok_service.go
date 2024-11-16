package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type TikTokService struct {
}

func NewTikTokService() *TikTokService {
	return &TikTokService{}
}

var tiktokDefaultParams = map[string]string{
	"aid":              "1988",
	"app_name":         "tiktok_web",
	"device_platform":  "web",
	"referer":          "",
	"user_agent":       "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0",
	"cookie_enabled":   "true",
	"screen_width":     "1920",
	"screen_height":    "1080",
	"browser_language": "en-US",
	"browser_platform": "Linux+x86_64",
	"browser_name":     "Mozilla",
	"browser_version":  "5.0+(X11)",
	"browser_online":   "true",
	"timezone_name":    "Asia/Jakarta",
	"priority_region":  "ID",

	"appId":   "1180",
	"region":  "ID",
	"appType": "t",

	"isAndroid":     "false",
	"isMobile":      "false",
	"isIOS":         "false",
	"OS":            "linux",
	"tt-web-region": "ID",
	"language":      "en-US",
	"verifyFp":      "verify_kjf974fd_y7bupmR0_3uRm_43kF_Awde_8K95qt0GcpBk",
}

func (s *TikTokService) GetFYPVideos(count int) (*string, error) {
	uri, err := url.Parse(viper.GetString("tiktok.base_url.fyp"))
	if err != nil {
		return nil, err
	}

	requestParams := map[string]string{
		"id":        "1",
		"from_page": "fyp",
		"count":     fmt.Sprint(count),
	}

	values := uri.Query()
	for key, value := range requestParams {
		values.Set(key, value)
	}

	uri.RawPath = values.Encode()

	log.Println(uri.String())

	resp, err := http.Get(uri.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	return nil, nil
}

func (s *TikTokService) SearchVideos(keyword string) {

}
