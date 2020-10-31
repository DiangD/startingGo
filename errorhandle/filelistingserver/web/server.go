package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"shmiloveu.fun/startingGo/errorhandle/filelistingserver/web/handle"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//统一的错误处理逻辑 -> 函数式编程
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//对panic做处理
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			//处理自定义错误
			if err, ok := err.(userError); ok {
				log.Print(err.Message())
				http.Error(writer, err.Message(), http.StatusInternalServerError)
				return
			}

			//处理常规错误
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			log.Print(err.Error())
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//自定义错误结构
type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(handle.FileList))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
