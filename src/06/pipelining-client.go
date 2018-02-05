package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// 今回は簡単化去せるために、予め送信メッセージを配列にためておいて、自動で完了するスクリプトにしている
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var err error

	requests := make(chan *http.Request, len(sendMessages))
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Access %d\n", current)
	defer conn.Close()

	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest("GET", "http://localhost:8080?message="+sendMessages[i], nil)

		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}

		if err != nil {
			panic(err)
		}

		// リクエストを送信する
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		fmt.Println("send: ", sendMessages[i])
		// requestsに全て配列に格納
		requests <- request
	}
	close(requests)

	// 全て送ったリクエストのレスポンスを確認するためのリーダを作成する
	reader := bufio.NewReader(conn)

	for request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}

		// レスポンスを一件一件確認していく
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))
		if current == len(sendMessages) {
			break
		}
	}
}
