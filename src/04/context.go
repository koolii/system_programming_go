package main

import (
	"context"
	"fmt"
)

// 他にも終了時間を設定したりタイムアウトの期限を設定できる
// context.WithDeadline()やcontext.WithTimeout()もある

func main() {
	fmt.Println("start sub()")

	// 終了を受け取るための終了関数月コンテキスト
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		cancel()
	}()

	// 終了を待機
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
