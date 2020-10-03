package util

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, errMsg *RestErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errMsg.StatusCode)
	json.NewEncoder(w).Encode(errMsg)
}
