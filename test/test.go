package main

import (
	"fmt"
	blog "github.com/BigCodilo/BestLogFramework/loger"
	"github.com/gorilla/mux"
	"net/http"
)

var Loger blog.BestLog

func main(){
	var err error
	Loger = blog.BestLog{}
	Loger.OnInfo()
	//Loger.OnDebug()
	//err = Loger.SetDebugPath("ddeebug")
	if err != nil{
		fmt.Println(err)
	}

	defer Loger.CloseFiles()
	StartServer()
}

func StartServer(){
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		//fmt.Println(os.Getenv("BLOG_INFO_LEVEL"))
		Loger.PrintLogs("jopa", "chlen", 7324)
		Loger.Info("PisaPisa")
		Loger.Debug("PisaPisaDebug")
		w.Write([]byte("hello"))
	})
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}