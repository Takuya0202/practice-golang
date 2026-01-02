// mainパッケージ - プログラムのエントリーポイント
package main

import (
	"fmt"
	"my-app/lesson_packages/models" // modelsパッケージをインポート（go.modのmodule名 + ディレクトリパス）
	"my-app/lesson_packages/utils"  // utilsパッケージをインポート
)

func main() {
	fmt.Println("===== パッケージの練習 =====")
	fmt.Println()

	// ===== modelsパッケージの使用 =====
	fmt.Println("--- modelsパッケージの使用 ---")

	// 公開関数NewUserを使う（大文字で始まるので使える）
	user := models.NewUser(1, "山田太郎", "yamada@example.com")
	fmt.Println("ユーザーを作成しました")

	// 公開メソッドSetAgeを使う
	user.SetAge(25)
	fmt.Println("年齢を設定しました")

	// 公開メソッドGetInfoを使う
	fmt.Println(user.GetInfo())

	// 公開メソッドGetAgeを使う（非公開フィールドageにアクセスする方法）
	fmt.Println("年齢:", user.GetAge())

	// 公開フィールドに直接アクセスできる（大文字で始まる）
	user.Name = "山田花子"
	fmt.Println("名前を変更:", user.Name)

	// 非公開フィールドには直接アクセスできない（コンパイルエラー）
	// user.age = 30  // エラー！ageは非公開（小文字で始まる）

	// 公開関数ValidateUserEmailを使う
	isValid := models.ValidateUserEmail("test@example.com")
	fmt.Println("メールアドレスの検証:", isValid)

	// 非公開関数は呼び出せない（コンパイルエラー）
	// models.validateEmail("test@example.com")  // エラー！validateEmailは非公開

	fmt.Println()

	// ===== utilsパッケージの使用 =====
	fmt.Println("--- utilsパッケージの使用 ---")

	// 公開関数Greetを使う（大文字で始まるので使える）
	greeting := utils.Greet("太郎")
	fmt.Println(greeting)

	// 公開関数FormatNumberを使う
	formatted := utils.FormatNumber(123)
	fmt.Println(formatted)

	// 非公開関数は呼び出せない（コンパイルエラー）
	// utils.greet("太郎")  // エラー！greetは非公開（小文字で始まる）
	// utils.formatNumber(123)  // エラー！formatNumberは非公開

	// 公開関数から非公開関数を呼び出す例
	// PublicGreetは公開関数だが、内部で非公開関数greetを呼び出している
	result := utils.PublicGreet("花子")
	fmt.Println("公開関数経由:", result)

	// 公開関数CalculateSumを使う
	sum := utils.CalculateSum(1, 2, 3, 4, 5)
	fmt.Println("合計:", sum)

	// 公開関数GetAverageを使う（内部で非公開関数calculateAverageを呼び出している）
	average := utils.GetAverage(10, 20, 30)
	fmt.Printf("平均: %.2f\n", average)

	fmt.Println()
}
