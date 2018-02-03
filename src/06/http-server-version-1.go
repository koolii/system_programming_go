package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("[logger] Server is running at localhost:8080")

	// この無限ループでリクエストを送信した時に、どういう挙動になっている？
	for {
		fmt.Println("[logger] Start New Loop and wait until a request come.")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// HTTPリクエストを分解
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			// デバッグ用の関数で、分解したHTTPリクエストを出力してくれる
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))

			// HTTPレスポンスを生成
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			// HTTPレスポンスを返す
			response.Write(conn)
			conn.Close()
		}()

		fmt.Println("[logger] Finish this loop, but it hasn't handle response to a client yet")
	}
}
