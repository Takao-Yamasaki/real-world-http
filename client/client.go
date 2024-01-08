package main

import (
	"log"
	"net/http"
)

func main() {
	// HEADメソッドでヘッダーを取得
	resp, err := http.Get("http://server:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
}
