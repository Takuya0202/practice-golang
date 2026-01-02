package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func lessonStruct() {
	// structで構造体を定義できる

	var user User
	user.Name = "山田"
	user.Age = 20
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)

	// 宣言と初期値の同時
	user2 := User{
		Name: "山田2",
		Age:  22,
	}
	fmt.Println(user2)

	// 関数の引数に構造体を渡すと都度copyされるので、ポインタを使うといい
	user3 := User{
		Name: "あああ",
		Age:  33,
	}
	showName(&user3)
}

func showName(user *User) {
	fmt.Println(user.Name)
}
