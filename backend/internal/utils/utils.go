// some kind of utils for different tasks
// sometimes without docs
package utils

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/logisshtick/mono/internal/vars"
)

// simply serialize to json and write struct as http response with status code
func WriteJsonAndStatusInRespone[T any](w http.ResponseWriter, status int, j *T) {
	w.WriteHeader(status)
	jsn, _ := json.Marshal(*j)
	w.Write(jsn)
}

// if contentLen > maxlen
// start f func, write response with error and return true
func ValidateContentLenAndSendResponse[T any](w http.ResponseWriter, contentLen int64, j *T) bool {
	if contentLen > vars.MaxHttpBodyLen {
		// here u can see some kind of runtime reflection swag
		// that get only enjoyers and other kinds of cool kids
		field := reflect.ValueOf(*j).FieldByName("Err")
		if field.IsValid() {
			if field.CanSet() {
				if field.Kind() == reflect.String {
					field.SetString(vars.ErrBodyLenIsTooBig.Error())
				}
			}
		}

		w.WriteHeader(http.StatusRequestEntityTooLarge)
		jsn, _ := json.Marshal(j)
		w.Write(jsn)
		return true
	}
	return false
}
