package main

import (
	"fmt"
	"sort"
)

func lessonMap() {
	// stringのketとintのvalueを持つmap
	// var m map[string]int
	// このままの宣言ではnilなので、makeを使う
	m := make(map[string]int)
	// 要素の追加
	m["apple"] = 100
	m["banana"] = 200
	m["cherry"] = 300
	m["apple2"] = 400
	fmt.Println(m)

	// 下でも行ける
	m2 := map[string]int{}
	m2["apple"] = 100
	m2["banana"] = 200
	fmt.Println(m2)

	// 要素の削除 delete関数を使う
	delete(m2, "apple")
	fmt.Println(m2)

	// for rangeを使うことでloopできる。mapは順序を持たないから、順番は保証されない。
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	// 配列、スライスのfor文 不使用の変数には_を使う
	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Println(v)
	}

	// mapには順序が存在しないため、キーを取り出して、ソートさせる。標準で順序付きmapはないみたい。
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key: %s, value: %d\n", k, m[k])
	}

	// 存在しないキー
	m3 := make(map[string]string)
	m3["apple"] = "apple"
	fmt.Println(m3["banana"]) // "" 存在しないキーは空文字。panicはしない。

	// キーが存在したかどうかは以下のようにして判断できる。これはgoは複数の戻り値を返却することができるから。
	// 第一引数はvalue、第二引数でキーが存在しているかどうかのboolを返却する。
	v, ok := m3["banana"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}
}
