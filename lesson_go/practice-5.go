package main

import "fmt"

// 配列とスライス
func lessonArrayAndSlice() {
	// 配列は固定長の要素
	var array [4]int
	array[0] = 1
	fmt.Println(array)

	// :=を使った配列宣言
	array2 := [4]int{}
	array2[0] = 1
	fmt.Println(array2)

	// 初期値を持った配列の宣言
	var array3 = [4]int{1, 2, 3}
	fmt.Println(array3) // [1 2 3 0]

	// スライスの宣言。nilとなる。
	// var slice []int
	// makeを使わないスライスはnilとなるため、アクセスできない。
	// slice[0] = 1
	// fmt.Println(slice)

	// makeをつかったスライス宣言。
	slice := make([]int, 4)
	fmt.Println(slice) //[0 , 0 , 0 , 0]

	// 初期値を持ったスライスの宣言
	slice2 := make([]int, 4, 8) //初期値(長さ)4、容量8
	fmt.Println(slice2)

	// 二次元の配列
	doubleArray := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(doubleArray)
	fmt.Println(doubleArray[0][0])

	// 二次元のスライス
	doubleSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(doubleSlice)
	fmt.Println(doubleSlice[0][0])

	// スライスの追加にはappendを使用
	doubleSlice = append(doubleSlice, []int{7, 8, 9, 10})
	fmt.Println(doubleSlice)

	// スライスの長さを取得
	fmt.Println("スライスの長さ(len)", len(doubleSlice))
	// スライスの容量を取得
	fmt.Println("スライスの容量(cap)", cap(doubleSlice))

	// 容量100,長さ0のスライス。容量を指定しておくことで、メモリの無駄を削減(アロケーションを抑える)
	newSlice := make([]int, 0, 100)
	fmt.Println(len(newSlice), cap(newSlice))

	// 1 ~ 100追加
	for i := 0; i < 100; i++ {
		newSlice = append(newSlice, i+1)
	}
	fmt.Println(newSlice)
	// スライスは添字でアクセスできる
	fmt.Println(newSlice[0:10]) // インデックス番号の0から9番目まで

	// スライスの削除。同じ型のスライスを作成して、詰める
	newSlice2 := make([]int, 0, len(newSlice)/2)
	for i := 0; i < len(newSlice); i++ {
		if i%2 != 0 {
			newSlice2 = append(newSlice2, newSlice[i])
		}
	}
	fmt.Println(newSlice2)

	// appendで削除
	n := 50
	// インデックス番号nの要素を削除
	// n-1番目の要素からn+1から最後までの要素を追加。スライスをそのまま渡すとエラーなので、...で展開
	// appendの第二引数はelem ...Typeとなっていて、スライスに対し、...とすることで数値を１つ１つ展開して渡すことができる。
	newSlice = append(newSlice[:n], newSlice[n+1:]...)
	fmt.Println(newSlice)

	// copy関数。copy元スライスとcopy先スライスを使用する
	// copy元スライスにはコピー分の長さが必要であり、copy関数の戻り値はcopyした要素数を返す。
	copySource := []int{1, 2, 3}
	copyTraget := make([]int, len(copySource))
	copyNum := copy(copyTraget, copySource)
	fmt.Println(copyNum)
	// copyで削除。
	m := 25
	// copy(newSlice[m:], newSlice[m+1:])でm+1番目から最後までの要素をm番目以降にコピー。
	newSlice = newSlice[:m+copy(newSlice[m:], newSlice[m+1:])]
	fmt.Println(newSlice)
}
