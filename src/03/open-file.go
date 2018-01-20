package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// io.Copyでファイル内容を全て標準出力に書き出す
	// これでコピーする方法がわかったから、別のファイルに読み込んだファイルをコピーしたかったらos.Stdoutを対象ファイルにすれば良さそう
	io.Copy(os.Stdout, file)
}
