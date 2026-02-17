package main

import "fmt"

// インターフェースはメソッドを持つ型のインターフェースを定義することができる
// メソッドとして、Error型を返すSpeak()を持つ場合、全てそれらはSpeakerインターフェースであるとみなされる。
type Speaker interface {
	Speak() error
}

type Dog struct{}

// implemntsなどで継承とかしてないけど、Speak()メソッドを持つので、Speakerインターフェースを満たす。
func (d *Dog) Speak() error {
	fmt.Println("わんわん")
	return nil
}

type Cat struct{}

func (c *Cat) Speak() error {
	fmt.Println("にゃーにゃー")
	return nil
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}
func lessonInterface2() {
	dog := Dog{}
	cat := Cat{}
	DoSpeak(&dog)
	DoSpeak(&cat)
}
