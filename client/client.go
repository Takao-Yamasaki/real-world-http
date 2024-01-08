package main

import (
	"log"
	"net/http"
	// "os"
	"strings"
)

// ファイルを渡す場合
// func main() {
// 	file, err := os.Open("client.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := http.Post("http://server:18888", "text/plain", file)
// 	if err != nil {
// 		// 送信失敗
// 		panic(err)
// 	}
// 	log.Println("Status:", resp.Status)
// }

// プログラム中で生成したテキストを渡す場合
func main() {
	reader := strings.NewReader("テキスト")
	resp, err := http.Post("http://server:18888", "text/plain", reader)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
