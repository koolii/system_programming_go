package main

import (
	"fmt"
	"math"
)

// 素数を求めて、チャネルを通じて返すコード

// チャネルを生成して、それをmain()に渡している
func primeNumber() chan int {
	result := make(chan int)

	go func() {
		result <- 2
		for i := 3; i < 100; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false

			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}

		close(result)
	}()

	return result
}

// このプログラムのように、計算しながら逐次返すようにして、その受け渡しにチャネルを使うと
// PythonやJavascriptのジェネレータのようなことを実現できる
func main() {
	pn := primeNumber()

	// ここがポイント
	// 帰ってきたチャネルは、for...range構文の中で配列と同じ場所に置くと、「値が来る度にforループが回る、個数が未定の動的配列」
	// のように扱うことができる。このforループは、チャネルがクローズされると止まる
	for n := range pn {
		fmt.Println(n)
	}
}
