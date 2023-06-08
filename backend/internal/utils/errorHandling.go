// some kind of utils for different tasks
// sometimes without docs
package utils

import (
	"net/http"
	"encoding/json"
	"reflect"
	"io"

	"github.com/logisshtick/mono/internal/vars"
)

// set error field in any struct if it exist
func SetErrorField[T any](j *T, err error) {
	// here u can see some kind of runtime reflection swag
	// that get only enjoyers and other kinds of cool kids
	s := reflect.ValueOf(j).Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName("Err")
		if f.IsValid() && f.CanSet() && f.Kind() == reflect.String {
			f.SetString(err.Error())
		}
	}
}

// if contentLen > maxlen
// start f func, write response with error and return true
func ErrWithContentLen[T any](w http.ResponseWriter, j *T, contentLen int64) bool {
	if contentLen > vars.MaxHttpBodyLen {
		return ErrNotNilSendResponse(w, j,
			http.StatusRequestEntityTooLarge,
			vars.ErrBodyLenIsTooBig,
		)
	}
	return false
}

// check error on request body reading
// if it exist send response with err
func ErrWithBodyReading[T any](w http.ResponseWriter, j *T, err error) bool {
	if err != nil {
		return ErrNotNilSendResponse(w, j,
			http.StatusInsufficientStorage,
			vars.ErrBodyReadingFailed,
		)
	}
	return false
}

// read request body and send 
// response if any error exist
func BodyReading[T any](w http.ResponseWriter, r *http.Request, j *T) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ErrNotNilSendResponse(w, &j,
			http.StatusInsufficientStorage,
			vars.ErrBodyReadingFailed,
		)
		return []byte{}, err
	}
	return body, nil
}

// unmarshal json and validate
func UnmarshalJson[T, Y any](w http.ResponseWriter, in *T, out *Y, bytes []byte) bool {
	err := json.Unmarshal(bytes, in)
	if err != nil {
		ErrNotNilSendResponse(w, out,
			http.StatusUnprocessableEntity,
			vars.ErrInputJsonIsIncorrect,
		)
		return true
	}
	return false
}

// return true and send http json respone and status code
func ErrNotNilSendResponse[T any](w http.ResponseWriter, j *T, status int, err error) bool {
	if err != nil {
		SetErrorField(j, err)
		WriteJsonAndStatusInRespone(w, j, status)
		return true
	}
	return false
}

