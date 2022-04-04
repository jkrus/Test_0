package app

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"wb_L0/internal/config"
	api "wb_L0/internal/errors"
	"wb_L0/internal/handlers"
	nats_subscr "wb_L0/internal/nats-subscr"
	"wb_L0/internal/services"
	"wb_L0/pkg/datastore/cache"
	"wb_L0/pkg/datastore/postgres"
	"wb_L0/pkg/server"
)

func startCommand(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "start",
		Usage: "Starts " + app.Usage,
		Before: func(context *cli.Context) error {
			// load data from config file.
			if err := cfg.Load(); err != nil {
				return api.ErrLoadConfig(err)
			}

			return nil
		},
		Action: func(context *cli.Context) error {
			return provideServices(ctx, wg, cfg)
		},
		After: func(c *cli.Context) error {
			<-c.Done()

			ctx.Done()
			wg.Wait()
			log.Println("Application shutdown complete.")

			return nil
		},
	})
}

// provideServices provides cli command specific services from application.
func provideServices(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config) error {
	orm, err := postgres.Start(ctx, cfg, wg)
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return api.ErrOpenDatabase(err)

		}
	}

	// provide chi router interface
	ri := chi.NewRouter()

	// provide HTTP server interface.
	s := server.NewHTTP(ctx, wg, cfg, ri)

	// provide memory cache.
	memoryCache := cache.NewCache(cfg)
	service, err := services.NewServices(ctx, wg, cfg, orm, memoryCache)
	if err != nil {
		return err
	}

	// provide NatsSubscribe Service.
	ns := nats_subscr.NewNatsSubscribeService(ctx, wg, cfg, orm, memoryCache)
	if err = ns.Start(); err != nil {
		return api.ErrStartNatsSubscribeService(err)
	}

	// provide Handlers.
	handlers.NewHandlers(service, ri).Register()

	if err = s.Start(); err != nil {
		if !errors.Is(http.ErrServerClosed, err) {
			return api.ErrStartHTTPServer(err)
		}
	}

	return nil
}
