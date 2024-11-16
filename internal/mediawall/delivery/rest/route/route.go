package route

import (
	"net/http"

	delivery "github.com/agilistikmal/media-wall-go/internal/mediawall/delivery/rest"
)

type Route struct {
	Mux *http.ServeMux

	TikTokHandler *delivery.TikTokHandler
}

func NewRoutes(tiktokHandler *delivery.TikTokHandler) *Route {
	return &Route{
		Mux:           http.NewServeMux(),
		TikTokHandler: tiktokHandler,
	}
}

func (r *Route) Init() {
	r.ProductRoutes()
}

func (r *Route) ProductRoutes() {
}
