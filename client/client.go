package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	// x-www-form-urlencoded形式のPOSTメソッド
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm("http://server:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
}
