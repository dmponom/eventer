package cli

import (
	"context"
	"go.uber.org/fx"
	"log"
)

func HandleFxApp(app *fx.App) error {
	startCtx, startCancel := context.WithTimeout(context.Background(), app.StartTimeout())
	defer startCancel()

	if err := app.Start(startCtx); err != nil {
		return err
	}

	sigs := app.Done()

	log.Printf("\n\n Received signal: %+v\n\n", <-sigs)
	log.Printf("Existing in %s\n", app.StartTimeout().String())

	stopCtx, stopCancel := context.WithTimeout(context.Background(), app.StartTimeout())
	defer stopCancel()

	log.Printf("Stopping app...\n")
	return app.Stop(stopCtx)
}
