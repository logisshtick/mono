// /login endpoint
//
// json example
//
// input:
//
//		{
//		  "email": "example@mail.com",
//		  "password": "example_password",
//	       "device_id": "any uniq to session key",
//		}
//
//		notes:
//		  max username len 64
//		  max password len 64
//		  max email    len 256
//
// output:
//
//	{
//	  "error": "nil",
//	  "access_token": 75298352711,
//	  "refresh_token": "cbe02bcc-6126-4079-8f2f-05a1562df0b0",
//	}
//
//	notes:
//	  if error is not "nil"
//	  u shouldnt check any other fields
//	  coz its can be incorrect or have any crap.
//	  access token type is uint64
//
// future fixes:
//   add util func for auto json validation as we want
//   add work with db
//   remove code duplications
package login

import (
	//"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/logisshtick/mono/internal/auth"
	"github.com/logisshtick/mono/internal/utils"
	"github.com/logisshtick/mono/internal/vars"
)

var (
	endPoint string
	mlog     *log.Logger
	wlog     *log.Logger
	elog     *log.Logger
)

type input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	DeviceId string `json:"device_id"`
}

type output struct {
	Err          string `json:"error"`
	AccessToken  uint64 `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Start(ep string, m *log.Logger, w *log.Logger, e *log.Logger) error {
	endPoint = ep
	mlog = m
	wlog = w
	elog = e

	err := auth.Init()
	if err != nil {
		return err
	}

	mlog.Printf("%s Login module inited\n", endPoint)
	return nil
}

// http handler that used as callback for http.HandleFunc()
func Handler(w http.ResponseWriter, r *http.Request) {
	mlog.Printf("%s connected %s\n", r.URL.Path, r.RemoteAddr)

	if utils.ValidateContentLenAndSendResponse(w, r.ContentLength, &output{}) {
		elog.Printf("%s body len too big: %d\n", r.URL.Path, r.ContentLength)
		return
	}

	input := input{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		elog.Printf("%s body reading failed: %v\n", r.URL.Path, err)
		utils.WriteJsonAndStatusInRespone(
			w, http.StatusInsufficientStorage,
			&output{
				Err: vars.ErrBodyReadingFailed.Error(),
			},
		)
		return
	}
	err = json.Unmarshal(body, &input)
	if err != nil {
		elog.Printf("%s body json is incorrect: %v\n", r.URL.Path, err)
		utils.WriteJsonAndStatusInRespone(
			w, http.StatusUnprocessableEntity,
			&output{
				Err: vars.ErrInputJsonIsIncorrect.Error(),
			},
		)
		return
	}

	const userId = 123
	deviceIdHash, err := auth.HashDeviceId(input.DeviceId)
	if err != nil {
		elog.Printf("%s deviceId hashing failed: %v\n", r.URL.Path, err)
		utils.WriteJsonAndStatusInRespone(
			w, http.StatusUnprocessableEntity,
			&output{
				Err: err.Error(),
			},
		)
		return
	}

	accessToken, refreshToken, err := auth.GenTokensPair(
		userId,
		deviceIdHash,
	)
	output := output{
		Err:          "nil",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	utils.WriteJsonAndStatusInRespone(
		w, http.StatusOK,
		&output,
	)
}

func Stop() error {
	mlog.Printf("%s Login module stoped\n", endPoint)
	return nil
}
