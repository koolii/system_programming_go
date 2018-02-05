package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for sessionResponse := range sessionResponses {
		// wait for finishing process
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))
	content := "Hello World\n"

	// write into response
	// keep up session
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          ioutil.NopCloser(strings.NewReader(content)),
	}

	// when it finishes process, write it into go-channel
	// ブロックされていた writeToConn() の処理を再始動する
	resultReceiver <- response
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	// セッション内のリクエストを順に処理するためのチャネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	// レスポンスを直列化してソケットに書き出す専用のgoroutine
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)

	for {
		// レスポンスを受け取ってセッションのキューに入れる
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読み込み
		request, err := http.ReadRequest(reader)

		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("[logger] Timeout")
				break
			} else if err == io.EOF {
				fmt.Println("[logger] EOF")
				break
			}
			panic(err)
		}

		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse

		// 非同期でレスポンスを実行
		go handleRequest(request, sessionResponse)
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
