package main

import "fmt"

// 小文字で始まる関数は同一パッケージ内からしか呼び出せない。
func lessonVariable() {
	// 変数の宣言。変数を使用しないとコンパイルエラーになる。
	var n int = 1 // 型を指定して変数を宣言。
	fmt.Println(n)

	// 型推論を使って、変数を定義
	a := 2   // int
	b := 1.5 // float64

	// 型が異なるため、コンパイルエラーになる。
	// result := a + b

	// 型を返還
	result := float64(a) + b
	fmt.Println(result)

	// 文字列の結合
	message := "Hello, " + "World!"
	fmt.Println(message)

	// 定数の宣言。未使用でもコンパイルエラーにならない。
	const x = 1
	y := 2

	// constでは使われ方によって型が決定され、untyped constantとなる
	f := 1.2 + (x + 2) + float64(y)
	fmt.Println(f)
}
