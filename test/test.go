package main

import (
	blog "github.com/BigCodilo/BestLogFramework/loger"
	"github.com/gorilla/mux"
	"net/http"
)

var Loger blog.BestLog

func main(){


	Loger = blog.NewBestLog()
	Loger.Info.TurnOn()
	Loger.Fatal.TurnOn()
	Loger.Fatal.SetFilePath("FatalFile")
	Loger.Debug.TurnOn()
	Loger.Debug.SetFilePath("debugishe")
	defer Loger.CloseFiles()
	StartServer()
}

func StartServer(){
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		//fmt.Println(os.Getenv("BLOG_INFO_LEVEL"))
		Loger.PrintAll("jopa", "chlen", 7324)
		Loger.Info.Print("PisaPisa")
		Loger.Debug.Print("NewTestDebug")
		w.Write([]byte("hello"))
	})
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}