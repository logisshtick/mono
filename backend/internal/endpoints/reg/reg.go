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

	"github.com/logisshtick/mono/internal/constant"
	"github.com/logisshtick/mono/internal/db"
	"github.com/logisshtick/mono/internal/utils"
	"github.com/logisshtick/mono/internal/vars"

	"github.com/cespare/xxhash"
	"github.com/jackc/pgx/v5"
	"github.com/logisshtick/mono/pkg/cryptograph"
	"github.com/logisshtick/mono/pkg/limiter"
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

	dbConn    *pgx.Conn
	limit     *limiter.Limiter[uint64]
	globalPow string
)

func Start(ep string, m *log.Logger, w *log.Logger, e *log.Logger) error {
	endPoint = ep
	mlog = m
	wlog = w
	elog = e

	var err error
	dbConn, err = pgx.Connect(context.Background(), utils.GetDbUrl())
	if err != nil {
		elog.Printf("%s database connection failed: %v", endPoint, err)
		return err
	}

	limit = limiter.New[uint64](10, 1800, 2048, 4096, 20)
	globalPow, err = utils.GetGlobalPow()
	if err != nil {
		elog.Printf("%s pow error: %v", endPoint, err)
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

	if !limit.Try(
		xxhash.Sum64String(r.UserAgent()),
	) {
		mlog.Printf("%s %s action limited", endPoint, r.RemoteAddr)
		out.Err = vars.ErrActionLimited.Error()
		utils.WriteJsonAndStatusInRespone(w, &out,
			http.StatusServiceUnavailable,
		)
		return
	}

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

	if len(in.Pass) > constant.C.RegMaxPasswordLen ||
		len(in.Nick) > constant.C.RegMaxNicknameLen ||
		len(in.Email) > constant.C.RegMaxEmailLen {
		elog.Printf("%s json fields is too big\n", endPoint)
		out.Err = vars.ErrFieldTooBig.Error()
		utils.WriteJsonAndStatusInRespone(w, &out,
			http.StatusUnprocessableEntity,
		)
		return
	}

	pow, err := cryptograph.GenRandPow(constant.C.MaxPowLen)
	if err != nil {
		elog.Printf("%s error with pow gen: %v", endPoint, err)
		out.Err = vars.ErrWithPowGen.Error()
		utils.WriteJsonAndStatusInRespone(w, &out,
			http.StatusInternalServerError,
		)
		return
	}

	hashedPass := cryptograph.HashPass(
		utils.PowCat(
			utils.PowCat(in.Pass, pow),
			globalPow,
		),
	)
	err = db.InsertUser(dbConn, in.Email, in.Nick, hashedPass, pow)
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
	if dbConn != nil {
		dbConn.Close(context.Background())
	}
	mlog.Printf("%s test module stoped\n", endPoint)
	return nil
}
