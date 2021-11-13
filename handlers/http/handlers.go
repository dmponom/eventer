package httphandlers

import (
	httpServer "eventer/core/http-server"
	"eventer/core/validators"

	"eventer/core/logger"
	"go.uber.org/fx"
)

type HTTPMethods struct {
	log       logger.Logger
	validator validators.Validator
}

func MakeHTTPHandlers(
	log logger.Logger,
	validator validators.Validator,
) httpServer.API {
	return &HTTPMethods{
		log:       log,
		validator: validator,
	}
}

var Module = fx.Options(
	fx.Provide(MakeHTTPHandlers),
)
