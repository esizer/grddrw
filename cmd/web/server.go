package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/esizer/grddrw/assets"
	"github.com/esizer/grddrw/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) serve() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/", app.devDisableCache(http.FileServerFS(assets.Files))))
	r.Get("/", templ.Handler(templates.Home()).ServeHTTP)
	r.Get("/grid", templ.Handler(templates.DrawPage(app.grid, app.palette)).ServeHTTP)
	r.Put("/pixel", app.paintHandler)
	r.Put("/clear", app.clearHandler)
	r.Post("/add-color", app.addColorHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", 3000),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
