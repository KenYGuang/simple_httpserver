package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

func hello(w http.ResponseWriter,req *http.Request){
	//a := req.URL.Query()
	//val := a["flag"][0]

	req.ParseForm();
	flag := req.Form.Get("flag")
	ip := req.Form.Get("ip")

	fmt.Println(flag,ip)
	w.Write([]byte("Hello"))
}




func main(){


	http.HandleFunc("/hello",hello)
	err := http.ListenAndServe(":8182",nil)
	if err != nil{
		log.Error(err)
	}
}
