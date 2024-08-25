package main

import (
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/a-h/templ"
	"github.com/esizer/grddrw/data"
	"github.com/esizer/grddrw/templates/components"
	"github.com/esizer/grddrw/toast"
)

func trimHash(h string) string {
	_, r := utf8.DecodeRuneInString(h)
	return h[r:]
}

func hex2RGB(h string) (data.Pixel, error) {
	trimmed := trimHash(h)
	vals, err := strconv.ParseUint(trimmed, 16, 32)
	if err != nil {
		return data.Pixel{}, err
	}

	pixel := data.Pixel{
		R: uint8(vals >> 16),
		G: uint8((vals >> 8) & 0xFF),
		B: uint8(vals & 0xFF),
	}

	return pixel, nil
}

func (app *application) paintHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	hex := r.Form.Get("color")
	pixel, err := hex2RGB(hex)
	if err != nil {
		t := toast.Danger("could not read hex")
		t.Serve(w, r, http.StatusBadRequest)
		return
	}

	templ.Handler(components.Pixel(&pixel)).ServeHTTP(w, r)
	return
}

func (app *application) clearHandler(w http.ResponseWriter, r *http.Request) {
	grid := data.NewGrid()

	templ.Handler(components.Grid(grid)).ServeHTTP(w, r)
	return
}
