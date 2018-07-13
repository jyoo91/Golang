package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	addr := ":80"
	r := mux.NewRouter()
	r.HandleFunc("/testpage", blockPage)

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func blockPage(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if ip == "10.0.0.170" {
		w.WriteHeader(200)
		w.Write([]byte("Welcome \n\n"))

		ua := r.Header
		//요청 헤더

		for idx, _ := range ua {
			//요청 헤더는 map으로 되어있음

			w.Write([]byte(idx + " : "))
			//idx --> key 의미

			data := r.Header.Get(idx)
			//key의 value 가져오기

			w.Write([]byte(data + "\n"))
		}
	} else {
		fmt.Println("conneting.... " + ip)
		w.WriteHeader(500)
		w.Write([]byte("block connections"))
	}
}
