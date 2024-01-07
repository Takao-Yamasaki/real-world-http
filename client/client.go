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
}
