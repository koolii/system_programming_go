package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	// 今回は簡単化去せるために、予め送信メッセージを配列にためておいて、自動で完了するスクリプトにしている
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	// var conn net.Conn = nil
	conn, _ := net.Dial("tcp", "localhost:8080")

	// リトライ用にループで全体を囲う
	for {
		var err error

		if conn == nil {
			// Dialから行ってconnを初期化
			// ここでエラーが出てしまう src/06/http-client.go:xx:xx: conn declared and not used
			// conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Access %d\n", current)
		}

		// POSTで文字列を送るリクエストを作成
		fmt.Println("[logger] Make a request to server")
		request, err := http.NewRequest("POST", "http://localhost:8080", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}

		fmt.Println("[logger] Send a request to server")
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		fmt.Println("[logger] Get response from server")
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			// タイムアウトはここでエラーになるのでリトライする
			fmt.Println("[logger] Retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}

		fmt.Println("[logger] Output response information")
		fmt.Println(string(dump))

		// 全部送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}
	}

	conn.Close()
}
