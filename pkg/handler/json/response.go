package json

import (
	"encoding/json"
	"net/http"
)

// JSON creates JSON response, ensuring successful is set on the body.
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
