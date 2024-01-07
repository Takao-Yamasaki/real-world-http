package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// リクエストに含まれる情報をテキストとして画面出力
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Set-Cookie", "VISIT=TRUE")
// 	// レスポンスヘッダーにCookieが存在すれば、訪問したことがある人
// 	if _, ok := r.Header["Cookie"]; ok {
// 		fmt.Fprintf(w, "<html><body>2回目以降</body></html>\n")
// 	} else {
// 		fmt.Fprintf(w, "<html><body>初回訪問</body></html>\n")
// 	}
// }

// httpサーバーを初期化
func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	// 18888番を使用
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
