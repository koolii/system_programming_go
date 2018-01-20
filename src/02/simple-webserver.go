package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriterはWebサーバだとブラウザに対してメッセージを書き込むのに使う
	// :8080にアクセスしてきたリクエストに対して、 `http.ResponseWriter sample`という文字列を返す
	io.WriteString(w, "http.ResponseWriter sample")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
