package main

import (
	"bufio"
	"fmt"
	"os"
)

// 標準入力、標準出力について
func lessonStandardInput() {
	// 入力読み込み
	scanner := bufio.NewScanner(os.Stdin)
	// 1行読み込み
	scanner.Scan()
	fmt.Println(scanner.Text())
}
