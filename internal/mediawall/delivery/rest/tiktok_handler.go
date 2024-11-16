package delivery

import (
	"github.com/agilistikmal/media-wall-go/internal/mediawall/service"
)

type TikTokHandler struct {
	service *service.TikTokService
}

func NewTikTokRESTHandler(service *service.TikTokService) *TikTokHandler {
	return &TikTokHandler{
		service: service,
	}
}
