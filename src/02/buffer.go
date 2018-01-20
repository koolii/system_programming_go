package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 他言語だとStringIO(Ruby)などと呼ばれている
	var buffer bytes.Buffer
	// Writeメソッドで書き込まれた内容をバッファとして纏めておいて、あとで結果を処理する
	buffer.Write([]byte("bytes.Buffer example\n"))
	// WriteStringメソッドを使えば文字列をそのまま受け取ってくれるが、他の構造体にはない機能
	buffer.WriteString("buffer.WriteString example\n")
	// ↑のメソッドを他の構造体でも使えるようにキャストを不要にしたメソッドがある
	// なんか本を参考に書いたけど、エラーで怒られた(Go 1.9.2)
	// io.WriteString(buffer, "io.WriteString examle\n")
	fmt.Println(buffer.String())
}
