package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/esizer/grddrw/data"
)

type config struct {
	dev bool
}

type application struct {
	config  config
	grid    data.Grid
	palette *data.Palette
}

func main() {
	var cfg config
	flag.BoolVar(&cfg.dev, "dev", true, "Run the app in development mode")

	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	g := data.NewGrid()

	app := &application{
		config:  cfg,
		grid:    g,
		palette: &data.Palette{},
	}

	logger.Info("Starting server", "addr", ":3000")

	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
