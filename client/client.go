package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://server:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	// ステータスを表示
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	// ヘッダーを表示
	log.Println("Headers:", resp.Header)
	// 特定のヘッダーを取得
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
