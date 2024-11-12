//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/agilistikmal/media-wall-go/internal/mediawall/delivery/rest"
	"github.com/agilistikmal/media-wall-go/internal/mediawall/delivery/rest/route"
	"github.com/agilistikmal/media-wall-go/internal/mediawall/service"
	"github.com/google/wire"
)

func InjectRoute() *route.Route {
	wire.Build(

		service.NewTikTokService,
		rest.NewTikTokHandler,

		route.NewRoutes,
	)

	return nil
}
