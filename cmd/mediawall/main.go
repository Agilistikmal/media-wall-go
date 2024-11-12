package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agilistikmal/media-wall-go/injector"
	"github.com/agilistikmal/media-wall-go/internal/infrastructure/config"
)

func main() {
	config.NewConfig()

	route := injector.InjectRoute()
	route.Init()

	port := 3333
	log.Printf("Running on http://localhost:%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), route.Mux)
}
