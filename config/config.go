package config

import (
	"github.com/jinzhu/configor"
	"go.uber.org/fx"
)

type Config struct {
	Logger     Logger
	HTTPServer HTTPServer
}

func Make() (*Config, error) {
	var cfg Config

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&cfg, "config/default.json"); err != nil {
		return nil, nil
	}

	return &cfg, nil
}

var Module = fx.Option(
	fx.Provide(Make),
)
