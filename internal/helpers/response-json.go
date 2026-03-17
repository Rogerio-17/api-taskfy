package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func ResponseWithError(w http.ResponseWriter, statusCode int, message string) {
	fmt.Println(message)
	ResponseWithJSON(w, statusCode, map[string]string{"error": message})
}
