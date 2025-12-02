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
}
