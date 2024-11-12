package rest

import (
	"net/http"

	"github.com/agilistikmal/media-wall-go/internal/mediawall/service"
)

type TikTokHandler struct {
	service *service.TikTokService
}

func NewTikTokHandler(service *service.TikTokService) *TikTokHandler {
	return &TikTokHandler{
		service: service,
	}
}

func (h *TikTokHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	oAuthURL := h.service.GetOAuthURL()
	http.Redirect(w, r, oAuthURL, 200)
}
