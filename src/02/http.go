package main

import (
	"io"
	"net"
	"os"
)

func main() {
	// これでTCP接続ができているコネクションを取得できる(実際は`net.TCPConn`構造体)
	conn, err := net.Dial("tcp", "ascii.jp:80")

	if err != nil {
		panic(err)
	}

	// 多分だけど、connがHTTP(GET)リクエストした結果をconnに書き込んでいる
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	// connに出力されたHTML情報をos.Stdoutにコピーしている(これができるのは`net.Conn`のインタフェースが`io.Reader`インタフェースでもあるから)
	io.Copy(os.Stdout, conn)
}
