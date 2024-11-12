package route

import (
	"net/http"

	"github.com/agilistikmal/media-wall-go/internal/mediawall/delivery/rest"
)

type Route struct {
	Mux *http.ServeMux

	TikTokHandler *rest.TikTokHandler
}

func NewRoutes(tiktokHandler *rest.TikTokHandler) *Route {
	return &Route{
		Mux:           http.NewServeMux(),
		TikTokHandler: tiktokHandler,
	}
}

func (r *Route) Init() {
	r.ProductRoutes()
}

func (r *Route) ProductRoutes() {
	r.Mux.HandleFunc("GET /oauth", r.TikTokHandler.Authorize)
	r.Mux.HandleFunc("GET /oauth/callback", r.TikTokHandler.Callback)
	r.Mux.HandleFunc("GET /video", r.TikTokHandler.QueryVideos)
}
