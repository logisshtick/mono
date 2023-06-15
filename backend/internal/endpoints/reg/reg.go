/* 
  /reg endpoint
  
  json example

  input:
    {
      "email": "example@mail.com",
      "nickname": "example_nickname",
      "password": "example_password",
    }
    
    notes:
      max username len 64
      max password len 64
      max email    len 256

  output:
    {
      "error": "nil",
    }
*/
package reg

import (
	"fmt"
	"log"
	"net/http"
)

var (
	endPoint string
	mlog     *log.Logger
	wlog     *log.Logger
	elog     *log.Logger
)

func Start(ep string, m *log.Logger, w *log.Logger, e *log.Logger) error {
	endPoint = ep
	mlog = m
	wlog = w
	elog = e

	mlog.Printf("%s reg module inited\n", endPoint)
	return nil
}

// http handler that used as callback for http.HandleFunc()
func Handler(w http.ResponseWriter, r *http.Request) {
	mlog.Printf("%s connected %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Test Endpoint finded!!! %s", r.URL.Path)
}

func Stop() error {
	mlog.Printf("%s test module stoped\n", endPoint)
	return nil
}

