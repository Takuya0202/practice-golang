package main

import (
	"fmt"
	"time"
)

func lessonGoroutine() {
	msg := "ごルーチンは呼び出し時に引数がキャプチャされる。"
	// goでは引数が値渡しでcopyされるので、絶対に上書きされない。
	go sendMessage(msg)
	// 無記名関数の場合はどちらが優先されるかわからない
	// 無記名関数の場合は、あくまで引数として受け取ってはいないので、無記名関数ないのsendMessageが呼び出された時点のmsgを使う。
	go func() {
		sendMessage(msg)
	}()
	msg = "なのでここで変更しても出力は変わらない。"
	time.Sleep(1 * time.Second)
}

func sendMessage(msg string) {
	fmt.Println(msg)
}
