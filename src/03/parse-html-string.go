package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))

	// bufio.NewReaderでnet.Connオブジェクトをbufio.Readerでラップしたオブジェクトを取得
	// このbufio.Readerオブジェクトをhttp.ReadResponse()に渡すとhttp.Response構造体のオブジェクトが取得できる
	// これで、ヘッダーやボディを個別に取得でき、パースできたことが分かる
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)

	// HTTPレスポンスのボディを書き出す
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

	// HTTPレスポンスのヘッダーを書き出す
	fmt.Println(res.Header)
}
