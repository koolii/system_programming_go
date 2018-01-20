package main

import (
	"bufio"
	"os"
)

func main() {
	// bufioと言うモジュールを使う
	// 出力結果を一時的にためておいて、ある程度の分量毎にまとめて書き出すことができる
	// Flushを呼び出す毎に出力される(このメソッドを実行しないといつまで経っても出力されない)
	// Flushを自動で呼び出す場合は、バッファサイズを第二引数に指定してbufio.Writerを作成する
	// => バッファ付き出力
	buffer := bufio.NewWriter(os.Stdout /* , 1048 */)
	buffer.WriteString("bufio.Writer\n")
	buffer.Flush()

	buffer.WriteString("example\n")
	buffer.Flush()
}
