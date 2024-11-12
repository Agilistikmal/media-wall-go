package route

import (
	"net/http"

	"github.com/agilistikmal/media-wall-go/internal/mediawall/delivery/rest"
)

type Route struct {
	Mux *http.ServeMux

	ProductHandler *rest.TikTokHandler
}

func NewRoutes(tiktokHandler *rest.TikTokHandler) *Route {
	return &Route{
		Mux:            http.NewServeMux(),
		ProductHandler: tiktokHandler,
	}
}

func (r *Route) Init() {
	r.ProductRoutes()
}

func (r *Route) ProductRoutes() {
	r.Mux.HandleFunc("GET /oauth", r.ProductHandler.Authorize)
}
