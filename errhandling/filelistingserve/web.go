package main

import (
	"learn_go/errhandling/filelistingserve/filelisting"
	log2 "log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log2.Printf("Panic: %v", r)
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			log.Warn("Error handling request: %s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(
					writer,
					userErr.Message(),
					http.StatusBadRequest,
				)
				return
			}
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFunc))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
