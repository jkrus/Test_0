package app

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/urfave/cli/v2"

	"wb_L0/internal/config"
)

// NewApp returns an application.
func NewApp(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config) *cli.App {

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"ver"},
		Usage:   "print the app version",
	}

	// construct cli application.
	app := &cli.App{
		Name:    config.AppName,
		Usage:   config.AppUsage,
		Version: version(),
		ExitErrHandler: func(context *cli.Context, err error) {
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	// register commands into cli application
	initCommand(cfg, app)
	startCommand(ctx, wg, cfg, app)

	return app
}

// Start starts the application.
func Start(ctx context.Context, app *cli.App) error {
	return app.RunContext(ctx, os.Args)
}
