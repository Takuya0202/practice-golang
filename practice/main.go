package main

import (
	"errors"
	"fmt"
)

// 変数宣言。デフォ値は0
var n int = 10

// Name は名前を表す VO
type Name string

// NewName は検証を通った値だけを Name として返す
func NewName(s string) (Name, error) {
	if s == "" {
		return "", errors.New("名前が空です")
	}
	if len([]rune(s)) > 10 {
		return "", errors.New("名前は10文字以内です")
	}
	return Name(s), nil
}

func main() {
	fmt.Println("hello world")
	fmt.Println(n)

	// VO はコンストラクタ関数を通して作る
	userName, err := NewName("gopher")
	fmt.Println(userName, err) // gopher <nil>

	_, err = NewName("")
	fmt.Println(err) // 名前が空です

	// 型変換ならチェックを素通りできてしまう
	badName := Name("")
	fmt.Printf("%q\n", badName) // ""
}
