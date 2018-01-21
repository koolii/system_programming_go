package main

import "fmt"

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan bool)

	// closure(jsで言うところのasyncになる？)
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		done <- true
	}()

	// 終了を待機(jsで言うところのawait)
	<-done
	fmt.Println("all tasks are finished")
}
