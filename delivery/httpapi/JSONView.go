package httpapi

import (
	"encoding/json"
	"net/http"
)

func JSONView(w http.ResponseWriter, data interface{}, statusCode int) {
	view, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(view)
}
