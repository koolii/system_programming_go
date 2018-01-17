package main

import "fmt"

// インターフェースを定義
type Talker interface {
	Talk()
}

// 構造体を宣言
type Greeter struct {
	name string
}

// 構造体はTalkerインターフェースで定義されているメソッドを持っている
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	// 下記二行が実際にGoコンパイラががインターフェースとの互換性をチェックしている箇所
	// インターフェースを満たす構造体のポインタは代入可能(Greeter型構造体インスタンスを作成し、ポインタを代入
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
