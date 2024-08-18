package main

import (
	"flag"
	"log/slog"
	"os"
)

type config struct {
	dev bool
}

type application struct {
	config config
}

func main() {
	var cfg config
	flag.BoolVar(&cfg.dev, "dev", true, "Run the app in development mode")

	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
	}

	logger.Info("Starting server", "addr", ":3000")

	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
