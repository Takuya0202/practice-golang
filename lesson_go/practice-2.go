package main

import (
	"fmt"
)

// iotaと定数を宣言する際に固有の値を自動で割り当てることができる。
const (
	Apple = iota
	Orange
	Banana
)

// 以下のようにしてstatus codeとかを定義することができる。
const (
	Success     = 200
	Failed      = 400
	ServerError = 500
)

type Fruit int
type Animal int

// このように定数名をenumのように宣言することができる。
const (
	Peach Fruit = iota
	Grape
	Kiwi
)

const (
	Dog Animal = iota
	Cat
	Bird
)

func lessonIota() {
	var fruit Fruit = Peach
	fmt.Println(fruit)
	// 下はAnimal型ではないため、コンパイルエラーとなる。
	// fruit = Dog
}
