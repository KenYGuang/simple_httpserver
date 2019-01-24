package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

func main(){


	//for i := 0; i <10; i++ {
		//starttime :=time.Now()
		_, err := http.Get("http://127.0.0.1:8182/hello?flag=false&ip=127.0.0.1")
		if err != nil {
			log.Error(err)
			fmt.Println("123",err)
			return
		}
		//endtime := time.Now()

		//resp.Body.Close()
		//time.Sleep(time.Second*30)
		//timeDelay := []byte(endtime.Sub(starttime).String())
		//fmt.Println(timeDelay,endtime.Sub(starttime))
		//
		//fp_write, err := os.OpenFile("http_1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
		//if err != nil{
		//	log.Error(err)
		//}
		//defer fp_write.Close()
		//timeDelay = append(timeDelay, '\n')
		//
		//_, err = fp_write.Write(timeDelay)
		//if err != nil {
		//	log.Error(err)
		//}
//	}

}
