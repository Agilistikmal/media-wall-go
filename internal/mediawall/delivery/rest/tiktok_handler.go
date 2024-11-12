package rest

import (
	"net/http"

	"github.com/agilistikmal/media-wall-go/internal/mediawall/service"
	"github.com/agilistikmal/media-wall-go/internal/pkg"
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

func (h *TikTokHandler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	oAuthResponse, err := h.service.GetAccessToken(code)
	if err != nil {
		pkg.SendError(w, 500, err.Error())
		return
	}

	pkg.SendSuccess(w, oAuthResponse)
}
