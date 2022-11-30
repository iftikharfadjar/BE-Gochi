package httptools

import (
	"net/http"
)

const (
	contentType = "Content-Type"
	plainText   = "text/plain"
	json        = "application/json"
)

func SendJSON(status int, w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set(contentType, json)
	return w
}
