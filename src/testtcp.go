package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"net"
	"time"
)

func main () {

	starttime :=time.Now()
	conn, err := net.Dial("tcp","127.0.0.1:8000")
	if err !=nil {
		log.Error(err)

	}
	endtime := time.Now()


	//time.Sleep(time.Second*30)
	//timeDelay := []byte(endtime.Sub(starttime).String())
	fmt.Println(conn,err,endtime.Sub(starttime))
}
