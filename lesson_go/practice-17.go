package main

import (
	"fmt"
	"sync"
	"time"
)

func lessonGoroutine2() {
	// ごルーチンの終了はsyncを使う。
	var wg sync.WaitGroup
	// ここで、n個のごルーチンを待ち受けることができる。
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("ごルーチン1終了")
		// 中身はadd(-1)と同じ。
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("ごルーチン2終了")
		wg.Done()
	}()

	// waitは0になるまで待つ。
	wg.Wait()
}
