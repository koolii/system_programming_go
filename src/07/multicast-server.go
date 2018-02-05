package main

import (
	"fmt"
	"net"
	"time"
)

const interval = 10 * time.Second

// クライアント側がソケットオープンして待ち受け、サーバがデータを送信する
// TCPとは全くの逆となる
func main() {
	fmt.Println("Start tick server at 224.0.0.1:9999")
	conn, err := net.Dial("udp", "224.0.0.1:9999")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	start := time.Now()
	// 中身は10秒単位の端数を取り出しているだけ
	// time.Duration(time.Sleep()の引数)は、実体はintだが、Go言語だと暗黙型変換ができないので明示的にキャストしている
	wait := start.Round(interval).Add(interval).Sub(start)
	time.Sleep(wait)
	ticker := time.Tick(interval)

	// 決まった時間感覚で定期的にfor文を回す為にtime.Tick()を使っている
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
