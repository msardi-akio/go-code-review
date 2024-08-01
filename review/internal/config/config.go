package config

import (
	"coupon_service/internal/api"
	"log"

	"github.com/brumhard/alligotor"
)

type Config struct {
	API api.Config
}

func New() Config {
	cfg := Config{}
	// MS investigate usage of this library
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
