package main

import "net/http"

func (app *application) devDisableCache(next http.Handler) http.Handler {
	if !app.config.dev {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
