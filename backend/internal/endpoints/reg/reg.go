/*
/reg endpoint

json example

input:

	{
	  "email": "example@mail.com",
	  "nick": "example_nickname",
	  "pass": "example_password",
	}

	notes:
	  max username len 64
	  max password len 64
	  max email    len 256

output:

	{
	  "err": "nil",
	}
*/
package reg

import (
	"context"
	"log"
	"net/http"

	"github.com/logisshtick/mono/internal/db"
	"github.com/logisshtick/mono/internal/utils"
	"github.com/logisshtick/mono/internal/vars"

	"github.com/jackc/pgx/v5"
)

type input struct {
	Email string `json:"email"`
	Nick  string `json:"nick"`
	Pass  string `json:"pass"`
}

type output struct {
	Err string `json:"err"`
}

var (
	endPoint string
	mlog     *log.Logger
	wlog     *log.Logger
	elog     *log.Logger

	dbConn *pgx.Conn
)

func Start(ep string, m *log.Logger, w *log.Logger, e *log.Logger) error {
	endPoint = ep
	mlog = m
	wlog = w
	elog = e

	var err error
	dbConn, err = pgx.Connect(context.Background(), utils.GetDbUrl())
	if err != nil {
		elog.Printf("%s database connection failed %v", endPoint, err)
		return err
	}

	mlog.Printf("%s reg module inited\n", endPoint)
	return nil
}

// http handler that used as callback for http.HandleFunc()
func Handler(w http.ResponseWriter, r *http.Request) {
	mlog.Printf("%s connected %s\n", endPoint, r.RemoteAddr)
	defer r.Body.Close()

	in := input{}
	out := output{}

	if utils.ErrWithContentLen(w, &out, r.ContentLength) {
		elog.Printf("%s body len too big: %d\n", endPoint, r.ContentLength)
		return
	}

	body, err := utils.BodyReading(w, r, &out)
	if err != nil {
		elog.Printf("%s body reading failed: %v\n", endPoint, err)
		return
	}

	if utils.UnmarshalJson(w, &in, &out, body) {
		elog.Printf("%s body json is incorrect\n", endPoint)
		return
	}

	if len(in.Pass) > 64 ||
		len(in.Nick) > 64 ||
		len(in.Email) > 256 {
		elog.Printf("%s json fields is too big\n", endPoint)
		out.Err = vars.ErrFieldTooBig.Error()
		utils.WriteJsonAndStatusInRespone(w, &out,
			http.StatusUnprocessableEntity,
		)
		return
	}

	err = db.InsertUser(dbConn, in.Email, in.Nick, in.Pass)
	if err != nil {
		elog.Printf("%s error with db insert: %v", endPoint, err)
		out.Err = vars.ErrWithDb.Error()
		utils.WriteJsonAndStatusInRespone(w, &out,
			http.StatusInternalServerError,
		)
		return
	}

	out.Err = "nil"
	utils.WriteJsonAndStatusInRespone(w, &out,
		http.StatusOK,
	)
}

func Stop() error {
	dbConn.Close(context.Background())
	mlog.Printf("%s test module stoped\n", endPoint)
	return nil
}
