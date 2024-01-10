package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/k0kubun/pp/v3"
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

const (
	username = "user"
	password = "pass"
)

// MD5ハッシュを生成
func getMD5(text string) string {
	hash := md5.New()
	io.WriteString(hash, text)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Digest認証の検証
func validateDigestAuth(authHeader string) bool {
	// Authorizationヘッダーの検証
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Digest" {
		return false
	}

	// パラメータの分割
	params := make(map[string]string)
	for _, param := range strings.Split(parts[1], ",") {
		parts := strings.SplitN(param, "=", 2)
		if len(parts) == 2 {
			params[strings.TrimSpace(parts[0])] = strings.Trim(parts[1], `"`)
		}
	}

	// 必要なパラメータの確認
	if params["username"] != username {
		return false
	}

	HA1 := getMD5(fmt.Sprintf("%s:%s:%s", username, params["realm"], password))
	HA2 := getMD5(fmt.Sprintf("%s:%s", "GET", params["uri"]))
	response := getMD5(fmt.Sprintf("%s:%s:%s:%s:%s:%s", HA1, params["nonce"], params["nc"], params["cnonce"], params["qop"], HA2))
	return response == params["response"]
}

// Digest認証のハンドラ
func handlerDigest(w http.ResponseWriter, r *http.Request) {
	pp.Printf("URL: %s\n", r.URL.String())
	pp.Printf("Query: %v\n", r.URL.Query())
	pp.Printf("Proto: %s\n", r.Proto)
	pp.Printf("Method: %s\n", r.Method)
	pp.Printf("Header: %v\n", r.Header)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--body--\n%s\n", string(body))
	// Authorizationヘッダーがあるか確認
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !validateDigestAuth(authHeader) {
		w.Header().Set("WWW-Authenticate", `Digest realm="Secret Zone", nonce="TgLc25U2BQA=f510a2780473e18e6587be702c2e67fe2b04afd", algorithm=MD5, qop="auth"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// 認証が成功した場合
	fmt.Fprintf(w, "<html><body>secret page</body></html>\n")
}

// httpサーバーを初期化
func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/digest", handlerDigest)
	log.Println("start http listening :18888")
	// 18888番を使用
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
