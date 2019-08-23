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
	Loger.Error.TurnOn()
	Loger.Fatal.SetFilePath("FatalFile.csv")
	Loger.Debug.TurnOn()
	Loger.Debug.SetFilePath("debugishe")
	//Loger.CloseFiles()
	StartServer()
}

func StartServer(){
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		//fmt.Println(os.Getenv("BLOG_INFO_LEVEL"))
		Loger.PrintAll("jopa", "chlen", 7324)
		Loger.Info.Print("PisaPisa")
		Loger.Debug.Print("111111111")
		w.Write([]byte("hello"))
	})
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}