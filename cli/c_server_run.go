package cli

import (
	"eventer/config"
	httpServer "eventer/core/http-server"
	"eventer/core/logger"
	"eventer/core/validators"
	httpHandlers "eventer/handlers/http"

	"flag"
	"go.uber.org/fx"
)

type runSystemCommand struct {
	flagSet *flag.FlagSet
}

func MakeRunSystemCommand() *runSystemCommand {
	return &runSystemCommand{
		flagSet: flag.NewFlagSet("run-system", flag.ContinueOnError),
	}
}

func (cmd *runSystemCommand) Name() string {
	return cmd.flagSet.Name()
}

func (cmd *runSystemCommand) Init(args []string) error {
	return cmd.flagSet.Parse(args)
}

func (cmd *runSystemCommand) Run() error {
	return HandleFxApp(fx.New(
		fx.StartTimeout(defaultAppStartTimeout),
		fx.StopTimeout(defaultAppStopTimeout),
		fx.Options(
			validators.Module,
			config.Module,
			logger.Module,
			httpHandlers.Module,
			httpServer.Module,
		),
	))
}
