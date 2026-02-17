package main

import (
	"fmt"
	"sync"
)

func lessonGoroutine3() {
	// ごルーチンのforループには注意が必要
	var wg sync.WaitGroup
	// 下の方法だとごルーチン内のiが参照される頃にはループが終わって、9が複数出力される可能性がある。
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}()
	// }
	// wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// ここでvにiをコピーしておくことで、ごルーチン内のiが参照される頃にはループが終わって、9が複数出力される可能性がなくなる。
		v := i
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()

	n := 0
	wg.Add(3)
	// ここも同様に例えばごルーチン1が参照した時のnが100でプラスする際にごルーチン2が足してしまった場合、1の誤差が生じる。最大30000であり、29999になる可能性あり。
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			n++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			n++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			n++
		}
	}()

	wg.Wait()
	fmt.Println(n) // 出力 25656

	// mutexを使うことで、競合を防ぐことができる。
	var mu sync.Mutex
	n = 0
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			// ロックを取得してからnを増やし、ロックを解除する。
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println(n) // 出力 30000
}
