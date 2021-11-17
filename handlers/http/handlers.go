package httphandlers

import (
	httpServer "eventer/core/http-server"
	"eventer/core/logger"
	"eventer/core/validators"
	"go.uber.org/fx"
)

type methods struct {
	log       logger.Logger
	validator validators.Validator
}

func MakeHTTPHandlers(
	log logger.Logger,
	validator validators.Validator,
) httpServer.API {
	return &methods{
		log:       log,
		validator: validator,
	}
}

var Module = fx.Options(
	fx.Provide(MakeHTTPHandlers),
)
