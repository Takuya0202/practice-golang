package main

import "fmt"

// メソッドを定義する
// Typeで宣言した方にはメソッドを定義することができる。
type Value int

// メソッドは func (レシーバー) メソッド名 (引数) 戻り値で定義できる
// レシーバーは何の方に対してメソッドを定義するかの部分
func (v Value) Add(n Value) Value {
	return v + n
}

func lessonMethod() {
	var v Value = 1
	fmt.Println(v.Add(2))

	v2 := Value(2)
	fmt.Println(v2.Add(3))
}
