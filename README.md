# real-world-http
## HTTP/1.0 基本となる4つの要素
- メソッドとパス
- ヘッダ
- ボディ
- ステータス

## HTTP/0.9にできることを試す
- ウェルノウンポートは、既にシステムで使用中だったり、一般ユーザーの権限では、それらのポートは使用できないように制限されている
- HTTPのデフォルトポートには、80番を使うが、その代替として`8000`/`8080`/`8888`がよく使用される
- HTTP/0.9で実現されているのは、`ボディの受信`と`パス`のみ。文書をブラウザからリクエストして、サーバーがデータを返すWebの基本骨格は完成している
- curlコマンドの実行
`ホスト名`または`IPアドレス`と`ポート番号`を指定
サーバー側が受け取るのは、`/greeting`というパスのみ
```
curl --http1.0 http://localhost:18888/greeting
<html><body>hello</body></html>
```
- サーバー側のログ
```
GET /greeting HTTP/1.0
Host: localhost:18888
Accept: */*
User-Agent: curl/7.79.1
```
- URLに検索キーワードを付与してリクエスト送信
`--get --data-urlencode` URLの末尾にクエリを付与
```
curl --http1.0 --get --data-urlencode "search word" http://localhost:18888
<html><body>hello</body></html>
```
