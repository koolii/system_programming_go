# Go言語について

## インターフェース
* Javaのインターフェースと基本同じと考えて良い
* Goだとインターフェース名に `~er, ~or` とつけることが多いらしい
* メンバを宣言できない(そのために構造体を利用する)
* 構造体も他の言語同じようなもの(メンバもメソッドも持てる)
* `func` 予約語とメソッド名の間にレシーバーを置くと、構造体にメソッドを定義したことになる => `interface.go#L16`
* 構造体にインターフェースを実装する時に `implements` 等のキーワードは必要ない
  * 勝手にGoコンパイラが判定してくれて、互換性チェックを実際に行っている箇所は `interface.go#24,25`)
* 副作用のあるメソッドではレシーバーの型をポインタ型にする(ここでは `(g *Greeter)`)
  * ポインタ型といっても、アクセス方法は変わらないし、C言語よりも扱いやすい

=> `go run docs/src/interface.go`