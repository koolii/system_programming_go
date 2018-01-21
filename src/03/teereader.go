package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	a, _ := ioutil.ReadAll(teeReader)
	// なんでか両方の返り値を`_`にすると、エラーで怒られた(go version 1.9.2)
	a = []byte{}
	fmt.Println(a)

	// だが、バッファには残っている
	fmt.Println(buffer.String())
}
