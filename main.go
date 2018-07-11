package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	addr := ":80"
	r := mux.NewRouter()
	r.HandleFunc("/", indexCTRL)

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func indexCTRL(w http.ResponseWriter, r *http.Request) {
	ip, port, _ := net.SplitHostPort(r.RemoteAddr)
	w.Write([]byte("Hello, Your Connect Info is " + ip + ":" + port))
}
