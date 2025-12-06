package main

import "fmt"

// ループについて
func lessonLoop() {
	// for文 1つ目。jsのようなパターン
	for i := 0; i < 10; i++ {
		fmt.Println("Hello,world1", i)
	}

	// while文パターン
	count := 0
	for count < 10 {
		fmt.Println("Hello,world2", count)
		count++
	}

	// 無限ループ forのみで宣言することで無限ループ
	i := 0
	for {
		fmt.Println("Hello,world3", i)
		i++
		if i > 10 {
			break
		}
	}

}
