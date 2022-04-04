package main

import (
	"log"

	"wb_L0/cmd/main/app"
	"wb_L0/internal/config"
)

func main() {
	ctx := app.NewContext()
	wg := app.NewWaitGroup()
	cfg := config.NewConfig()

	err := app.Start(ctx, app.NewApp(ctx, wg, cfg))
	if err != nil {
		log.Fatal(err)
	}
}
