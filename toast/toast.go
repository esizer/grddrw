package toast

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	INFO    = "info"
	SUCCESS = "success"
	WARNING = "warning"
	DANGER  = "danger"
)

type Toast struct {
	Level   string
	Message string
}

func New(level string, message string) Toast {
	return Toast{level, message}
}

func Info(message string) Toast {
	return New(INFO, message)
}

func Success(message string) {
	New(SUCCESS, message)
}

func Warning(message string) Toast {
	return New(WARNING, message)
}

func Danger(message string) Toast {
	return New(DANGER, message)
}

func (t Toast) Error() string {
	return fmt.Sprintf("%s: %s", t.Level, t.Message)
}

func (t Toast) jsonify() (string, error) {
	t.Message = t.Error()
	eventMap := map[string]Toast{}
	eventMap["toast"] = t
	jsonData, err := json.Marshal(eventMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (t Toast) Serve(w http.ResponseWriter, r *http.Request, code int) {
	data, err := t.jsonify()
	if err != nil {
		w.Header().Set("HX-Trigger", "{\"toast\": {\"level\": \"danger\", \"message\": \"Error\"}}")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("HX-Trigger", data)
	w.WriteHeader(code)
}
