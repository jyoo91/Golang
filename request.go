package main

import (
	"fmt"
	"log"
	//	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux" //원격경로 패키지
	// go get github.com/gorilla/mux ---> 패키지 받아오기
)

//NewRouter returns a new router instance.
//func NewRouter() *Router {			-----> 리턴값이 포인터
//        return &Router{namedRoutes: make(map[string]*Route), KeepContext: false}
//}

func main() {
	addr := ":80"                       //80번 포트에서 웹서버 실행
	r := mux.NewRouter()                //동적 URL 경로 처리
	r.HandleFunc("/testpage", testPage) //HandleFunc : 경로별로 요청을 처리할 핸들러 함수를 등록

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe()) //HTTP 연결받고 요청에 응답
	// 에러발생하면 에러메시지 출력하고 빠져나감 (메시지를 출력하고 os.Exit(1)호출하여 프로그램 종료 )
}

func testPage(w http.ResponseWriter, r *http.Request) {
	//	ua := r.Header.Get("User-Agent")
	ua := r.Header
	for idx, _ := range ua {
		w.Write([]byte(idx + " : "))
		data := r.Header.Get(idx)
		w.Write([]byte(data + "\n"))
	}

	fmt.Println(ua)
}
