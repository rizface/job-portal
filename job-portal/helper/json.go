package helper

import (
	"encoding/json"
	"job-portal/app/model/response"
	"net/http"
)

func JsonWriter(w http.ResponseWriter, code int, status string, data map[string]interface{}) {
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response.Standard{
		Code:   code,
		Status: status,
		Data:   data,
	})
}
