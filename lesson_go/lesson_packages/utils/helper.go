// utilsパッケージ - ユーティリティ関数を定義するパッケージ
package utils

import "fmt"

// Greet - 公開関数（大文字で始まる）
// 他のパッケージから呼び出せる
// 名前を受け取って挨拶文を返す
func Greet(name string) string {
	return fmt.Sprintf("こんにちは、%sさん！", name)
}

// FormatNumber - 公開関数
// 数値をフォーマットして返す
// 他のパッケージから呼び出せる
func FormatNumber(num int) string {
	return fmt.Sprintf("数値: %d", num)
}

// greet - 非公開関数（小文字で始まる）
// このパッケージ内でのみ使える
// 他のパッケージからは呼び出せない
func greet(name string) string {
	return fmt.Sprintf("hello, %s", name)
}

// formatNumber - 非公開関数
// このパッケージ内でのみ使える
func formatNumber(num int) string {
	return fmt.Sprintf("number: %d", num)
}

// PublicGreet - 公開関数から非公開関数を呼び出す例
// 公開関数内では同じパッケージの非公開関数を呼び出せる
// これにより、内部実装を隠しつつ、公開APIを提供できる
func PublicGreet(name string) string {
	// 同じパッケージ内なので非公開関数greetを呼び出せる
	return greet(name)
}

// CalculateSum - 公開関数
// 複数の数値の合計を計算する
func CalculateSum(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// calculateAverage - 非公開関数
// 平均値を計算する（内部実装）
func calculateAverage(numbers ...int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// GetAverage - 公開関数
// 非公開関数calculateAverageを呼び出す公開API
func GetAverage(numbers ...int) float64 {
	return calculateAverage(numbers...)
}
