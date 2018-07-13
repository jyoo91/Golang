package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	addr := ":80"
	r := mux.NewRouter()
	r.HandleFunc("/testpage", testPage)
	//HandleFunc : 경로별로 요청을 처리할 핸들러 함수를 등록

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func testPage(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println(ua)
}
