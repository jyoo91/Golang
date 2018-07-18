package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fake_client() {
	// GET 호출
	resp, err := http.Get("http://www.monitorapp.com/kr/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return data
}
