package main

import (
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/esizer/grddrw/data"
	"github.com/esizer/grddrw/templates/components"
	"github.com/esizer/grddrw/toast"
)

func (app *application) addColorHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	hex := r.Form.Get("color")

	_, err := app.palette.Add(hex)
	if err != nil {
		t := toast.Danger(err.Error())
		if errors.Is(err, data.HexExists) {
			t = toast.Warning(err.Error())
		}
		t.Serve(w, r, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	templ.Handler(components.PaletteSwatches(app.palette)).ServeHTTP(w, r)
	return
}
