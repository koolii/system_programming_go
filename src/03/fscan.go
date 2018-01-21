package main

import (
	"fmt"
	"strings"
)

var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)

	var i int
	var f, g float64
	var s string
	// Fscanfという関数を使えばスペース区切りの制約だけじゃなく、任意の区切りにすることができ	る
	fmt.Fscan(reader, &i, &f, &g, &s)
	// Go言語は型情報をデータが持っているので%vと書くだけで変数の型を読み取って変換してくれる
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
