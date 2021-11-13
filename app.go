package main

import (
	"context"
	"eventer/cli"
	"eventer/core/logger"
	"os"
)

func main() {
	if err := cli.Execute(os.Args[1:]); err != nil {
		logger.MakeRaw("eventer").Error(context.Background(), err.Error())
		os.Exit(1)
	}
}
