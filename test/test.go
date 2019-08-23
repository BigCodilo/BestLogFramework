package main

import (
	blog "github.com/BigCodilo/BestLogFramework/loger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Loger blog.BestLog

func main(){
	Loger = blog.NewBestLog()
	Loger.Error.TurnOn()
	Loger.Error.TurnOnCache()
	Loger.Error.SetFilePath("ErrorFileForLoging")
	StartServer()
}

func StartServer(){
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		Loger.Error.PrintWithCache("jewdnkeiwnefjneoifnaliuaeflauhfluei")
		w.Write([]byte("hello"))
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

