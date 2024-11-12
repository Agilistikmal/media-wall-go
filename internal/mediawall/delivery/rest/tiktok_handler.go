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
	http.Redirect(w, r, oAuthURL, http.StatusSeeOther)
}

func (h *TikTokHandler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	oAuthResponse, err := h.service.GetAccessToken(code)
	if err != nil {
		pkg.SendError(w, 500, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "access_token",
		Value: oAuthResponse.AccessToken,
		Path:  "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: oAuthResponse.RefreshToken,
		Path:  "/",
	})

	pkg.SendSuccess(w, oAuthResponse)
}

func (h *TikTokHandler) QueryVideos(w http.ResponseWriter, r *http.Request) {
	accessToken, err := r.Cookie("access_token")
	if err != nil {
		pkg.SendError(w, 500, err.Error())
		return
	}

	queryVideos, err := h.service.QueryVideos(accessToken.Value)
	if err != nil {
		pkg.SendError(w, 500, err.Error())
		return
	}

	pkg.SendSuccess(w, queryVideos)
}
