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

	// map,スライスでイテレーションする
	// rangeを使うことで、マップやスライスを簡単にループできる

	// スライスのイテレーション
	// スライスを定義
	numbers := []int{10, 20, 30, 40}

	// スライス: インデックスと値の両方を取得
	for index, value := range numbers {
		// indexはインデックス（0, 1, 2, 3...）
		// valueは値（10, 20, 30, 40...）
		fmt.Println(index, value)
	}

	// スライス: インデックスのみ取得（値は不要な場合）
	for index := range numbers {
		// インデックスのみ取得
		fmt.Println(index)
	}

	// スライス: 値のみ取得（インデックスは不要な場合）
	// ブランク識別子_を使って、不要な値を無視する
	for _, value := range numbers {
		// 値のみ取得
		fmt.Println(value)
	}

	// マップ（辞書）のイテレーション
	// マップを定義（Pythonのdict.items()と同じような動作）
	fruits := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// マップ: キーと値の両方を取得
	for key, value := range fruits {
		// keyはキー（"apple", "banana", "cherry"）
		// valueは値（1, 2, 3）
		fmt.Println(key, value)
	}

	// マップ: キーのみ取得（値は不要な場合）
	for key := range fruits {
		// キーのみ取得
		fmt.Println(key)
	}

	// マップ: 値のみ取得（キーは不要な場合）
	// ブランク識別子_を使って、不要な値を無視する
	for _, value := range fruits {
		// 値のみ取得
		fmt.Println(value)
	}
}
