package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
	"github.com/esizer/grddrw/data"
	"github.com/esizer/grddrw/templates/components"
)

func readChan(f url.Values, k string) (uint8, error) {
	s := f.Get(k)
	i, err := strconv.ParseUint(s, 0, 8)
	if err != nil {
		return 0, err
	}

	return uint8(i), nil
}

func (app *application) paintHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ri, err := readChan(r.Form, "r")
	if err != nil {
		t := Danger("Could not read red channel")
		t.Write(w, r, http.StatusBadRequest)
		return
	}

	bi, err := readChan(r.Form, "b")
	if err != nil {
		t := Danger("Could not read blue channel")
		t.Write(w, r, http.StatusBadRequest)
		return
	}

	gi, err := readChan(r.Form, "g")
	if err != nil {
		t := Danger("Could not read green channel")
		t.Write(w, r, http.StatusBadRequest)
		return
	}

	p := data.Pixel{
		R: ri,
		G: gi,
		B: bi,
	}

	templ.Handler(components.Pixel(&p)).ServeHTTP(w, r)
	return
}
