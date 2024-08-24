package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/esizer/grddrw/templates/components"
)

const (
	DANGER = "danger"
)

type Toast struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

func NewToast(severity string, message string) Toast {
	return Toast{severity, message}
}

func Danger(message string) Toast {
	return NewToast(DANGER, message)
}

func (t Toast) Error() string {
	return fmt.Sprintf("%s: %s", t.Severity, t.Message)
}

func (t Toast) toJson() ([]byte, error) {
	t.Message = t.Error()

	event := map[string]Toast{}
	event["toast"] = t

	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t Toast) Write(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("HX-Reswap", "none")
	w.WriteHeader(status)

	templ.Handler(components.Error(t.Message)).ServeHTTP(w, r)

	return
}
