package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// http://www.aozora.gr.jp/cards/000121/card628.html
var contents = []string{
	"これは、私《わたし》が小さいときに、村の茂平《もへい》というおじいさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山《なかやま》というところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐《ぎつね》」という狐がいました。",
	"ごんは、一人《ひとり》ぼっちの小狐で、しだ［＃「しだ」に傍点］の一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	// Accept後のソケットで何度も応答を返す為にループ
	// ここでループを追加することによって、TCPコネクションが張られた後も何度もリクエストを処理出来るようになる
	// => 無限ループ
	for {
		fmt.Println("[logger] From now on, it sets timeout into conn")
		// HTTPリクエストを分解
		// net.Connをbufio.Readerでラップして、それをhttp.ReadRequest()関数に渡している
		request, err := http.ReadRequest(bufio.NewReader(conn))

		if err != nil {
			if err == io.EOF {
				fmt.Println("[logger] EOF")
				break
			}
			panic(err)
		}

		// デバッグ用の関数で、分解したHTTPリクエストを出力してくれる
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		// HTTPレスポンスを生成
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))

		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}

		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("[logger] Server is running at localhost:8080")

	for {
		// コネクションを生成する時はここからスタートする
		fmt.Println("[logger] Create a new TCP connection and wait until a request come.")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// リクエストの処理を一任
		go processSession(conn)

		// コネクションがタイムアウトすると、次のリクエストからは下記のログからスタートする
		fmt.Println("[logger] Finish this loop, but it hasn't handle response to a client yet")
	}
}
