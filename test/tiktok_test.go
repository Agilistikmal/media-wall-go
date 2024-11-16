package test

import (
	"log"
	"testing"

	"github.com/agilistikmal/media-wall-go/internal/infrastructure/config"
	"github.com/agilistikmal/media-wall-go/internal/mediawall/service"
)

func init() {
	config.NewConfig()
}

func TestFYP(t *testing.T) {
	s := service.NewTikTokService()

	_, err := s.GetFYPVideos(5)
	if err != nil {
		log.Fatal(err)
	}
}
