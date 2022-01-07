package helper

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetParamsValue(r *http.Request, key string) string {
	params := mux.Vars(r)
	return params[key]
}

func GetTokenValue(claim interface{}, key string) string {
	var data map[string]interface{}
	byteJson,_ := json.Marshal(claim)
	json.Unmarshal(byteJson,&data)
	if key == "id" {
		return data["Id"].(string)
	}
	return data[key].(string)
}