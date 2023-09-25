package httpformat

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	Respond(w, r, code, map[string]string{"error": err.Error()})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			logrus.Warn(err)
		}
	}
}
