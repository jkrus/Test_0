package app

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"wb_L0/internal/config"
)

// initCommand appends initialization action into application.
func initCommand(cfg *config.Config, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "init",
		Usage: "Initialize " + config.AppUsage,
		Action: func(context *cli.Context) error {
			if err := cfg.Init(); err != nil {
				return errors.Wrap(err, "config init")
			}

			return nil
		},
	})
}
