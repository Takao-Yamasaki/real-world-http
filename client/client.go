package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// multipart/form-data形式でのファイルの送信
func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	//　ファイル読み込み失敗
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	// ファイルの全てのコンテンツをコピー
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://server:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
