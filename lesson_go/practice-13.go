package main

import (
	"fmt"
	"reflect"
)

func lessonInterface() {
	// インターフェースはどの方でも格納することができる
	var i interface{}
	i = 1

	// 方を取り出すにはtype assertionを使う
	num := i.(int)
	fmt.Println(num)

	// 間違った型を取り出すとpanicになるので以下のように確認する
	num, ok := i.(int)
	if !ok {
		fmt.Println("文字列")
	} else {
		fmt.Println(num)
	}

	// 1行で書くこともできる
	// goにはif 初期ステートメント ; 条件式 という書き方ができる
	if s, ok := i.(string); !ok {
		fmt.Println("文字列")
	} else {
		fmt.Println(s)
	}

	integer := 1
	PrintDetail(integer)
	type V int
	// したはV型なので知らない型
	v := V(1)
	PrintDetail(v)
	// reflectを使うと知らない型に対してもtype assertionを使うことができる
	PrintDetail2(v)
	PrintDetail3(v)
}

// 下によれば、いろんな方に対応できる
func PrintDetail(i interface{}) {
	switch t := i.(type) {
	case int, float32, float64:
		fmt.Println("数値", t)
	case string:
		fmt.Println("文字列", t)
	default:
		fmt.Println("その他", t)
	}
}

// reflrectパッケージを使うことで独自の型に対してもtype assertionを使うことができる
func PrintDetail2(i interface{}) {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Int, reflect.Float32, reflect.Float64:
		fmt.Println("数値", t)
	case reflect.String:
		fmt.Println("文字列", t)
	default:
		fmt.Println("その他", t)
	}
}

// interfaceはanyに置き換えることもできる
func PrintDetail3(i any) {
	switch t := i.(type) {
	case int, float32, float64:
		fmt.Println("数値", t)
	case string:
		fmt.Println("文字列", t)
	default:
		fmt.Println("その他", t)
	}
}
