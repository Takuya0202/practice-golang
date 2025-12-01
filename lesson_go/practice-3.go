package main

import (
	"fmt"
)

// 条件分岐について
func lessonIf() {
	x := 10
	// 条件分に()はいらない。
	if x > 5 && x < 15 {
		fmt.Println("xは5より大きく15より小さい")
	} else {
		fmt.Println("xは5以下か15以上")
	}

	// switch文 {}で囲って、:で処理
	switch x {
	case 10:
		fmt.Println("xは10")
	default:
		fmt.Println("わからん")
	}

	// 式を書くこともできる
	switch {
	case x > 5 && x < 15:
		fmt.Println("xは5より大きく15より小さい")
	default:
		fmt.Println("わからん")
	}
}
