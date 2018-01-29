package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		// 1リクエスト処理中に他のリクエストのAccept()が行えるように
		// goroutineで非同期にレスポンスを処理する
		go func() {
			fmt.Println(conn)
			// connを使った読み書き
		}()
	}
}
