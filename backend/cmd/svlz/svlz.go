package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "github.com/logisshtick/mono/internal/route"
	"github.com/logisshtick/mono/internal/test"
)

type endPoint struct {
	name    string
	start   func(string, *log.Logger, *log.Logger, *log.Logger) error
	handler func(http.ResponseWriter, *http.Request)
	stop    func() error
}

var (
	endPoints = [...]endPoint{
		{"/swag", test.Start, test.Handler, test.Stop},
		// {"/route", route.Start, route.Handler, route.Stop},
	}

	mlog *log.Logger
	wlog *log.Logger
	elog *log.Logger

	flagLogFile        string
	flagPort           string
	flagEndPointPrefix string
)

func endPointFailure(err error) {
	if err != nil {
		elog.Fatal(err)
	}
}

func init() {
	flag.StringVar(&flagLogFile, "log-file", "NULL", "set log file location")
	flag.StringVar(&flagPort, "port", "4433", "set http server port")
	flag.StringVar(&flagEndPointPrefix, "prefix", "/api", "set prefix for endpoints")
	flag.Parse()

	if flagLogFile == "NULL" {
		mlog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		wlog = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
		elog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	} else {
		if strings.HasSuffix(flagLogFile, "GEN") {
			flagLogFile = flagLogFile[:len(flagLogFile)-3] + time.Now().Format("2006-01-02 15-04-05") + ".log"
		}
		file, err := os.OpenFile(flagLogFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0666)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		mlog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		wlog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
		elog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	}

	for _, e := range endPoints {
		err := e.start(e.name, mlog, wlog, elog)
		endPointFailure(err)
	}
}

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-signalChannel
		switch s {
		case syscall.SIGHUP:
			elog.Println("SIGHUP received")
			stop()
			os.Exit(0)
		case syscall.SIGINT:
			elog.Println("SIGINT reveived")
			stop()
			os.Exit(0)
		case syscall.SIGTERM:
			elog.Println("SIGTERM received")
			stop()
			os.Exit(0)
		case syscall.SIGQUIT:
			elog.Println("SIGQUIT received")
			stop()
			os.Exit(0)
		default:
			elog.Println("unknown SIG received")
			stop()
			os.Exit(1)
		}
	}()

	for _, e := range endPoints {
		http.HandleFunc(flagEndPointPrefix+e.name, e.handler)
	}

	err := http.ListenAndServe(":"+flagPort, nil)
	if err != nil {
		stop()
		elog.Fatal(err)
	}
}

func stop() {
	for _, e := range endPoints {
		err := e.stop()
		if err != nil {
			elog.Fatal(err)
		}
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
