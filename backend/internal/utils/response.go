package utils

import (
	"encoding/json"
	"net/http"
)

// simply serialize to json and write struct as http response with status code
func WriteJsonAndStatusInRespone[T any](w http.ResponseWriter, j *T, status int) {
	w.WriteHeader(status)
	jsn, _ := json.Marshal(*j)
	w.Write(jsn)
}
