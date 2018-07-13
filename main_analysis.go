package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux" // 원격경로 패키지
	// go get github.com/gorilla/mux ---> 패키지 받아오기
)

func main() {
	addr := ":80"                // 80번 포트에서 웹서버 실행
	r := mux.NewRouter()         // 동적 URL 경로 처리
	r.HandleFunc("/", indexCTRL) // HandleFunc : 경로별로 요청을 처리할 핸들러 함수를 등록

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second, // 연결 허용 ~ 요청 본문 읽는 시점까지 걸리는 시간
		ReadTimeout:  15 * time.Second, // 요청 헤더 읽기 ~ 응답 종료까지 걸리는 시간
		// 신뢰할 수 없는 클라이언트와 네트워크일 때 연결 유지 못하도록 시간 설정
	}
	log.Fatal(server.ListenAndServe())
	//HTTP 연결받고 요청에 응답
	// 에러발생하면 에러메시지 출력하고 빠져나감 (메시지를 출력하고 os.Exit(1)호출하여 프로그램 종료 )
}

func indexCTRL(w http.ResponseWriter, r *http.Request) {
	// 경로에 접속했을 때 실행할 함수
	// (HTTP Response에 무언가를 쓸수 있게함, 입력된 request 요청 검토)

	ip, port, _ := net.SplitHostPort(r.RemoteAddr)
	// net.SplitHostPort : 호스트이름과 포트번호 포함한 문자열 반환

	w.Write([]byte("Hello, Your Connect Info is " + ip + ":" + port))
}
