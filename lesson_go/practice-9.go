package main

import (
	"fmt"
	"strings"
)

// ===== 型エイリアスにメソッドを追加する例 =====

// Email型を定義（stringのエイリアス）
// これにより、string型にEmail専用のメソッドを追加できる
type Email string

// Email型のメソッドを定義
// func (レシーバ 型) メソッド名() 戻り値 の形式
// (e Email) の部分がレシーバ（このメソッドがどの型に属するかを示す）
func (e Email) IsValid() bool {
	// メールアドレスの形式チェック
	// string(e) でEmail型をstring型に変換してから、strings.Containsで"@"が含まれているかチェック
	return strings.Contains(string(e), "@")
}

// Age型を定義（intのエイリアス）
type Age int

// Age型のメソッドを定義
func (a Age) IsAdult() bool {
	// 年齢が18歳以上かどうかを判定
	// aはAge型だが、intとして比較できる（型エイリアスは元の型と同じように使える）
	return a >= 18
}

func lessonType() {
	// typeで型エイリアスを作成できる。
	type MyString string
	var s MyString = "hello"
	fmt.Println(s)

	// stringに変換する場合
	s2 := string(s)
	fmt.Println(s2)

	// ===== 使い方 =====

	// Email型の変数を作成
	// Email("user@example.com") でstringをEmail型に変換
	email := Email("user@example.com")

	// Email型のメソッドIsValid()を呼び出す
	// email.IsValid() で、このemail変数に対してIsValidメソッドを実行
	// ここがメソッド。Email型なのでisValidメソッドを呼び出せる。
	if email.IsValid() {
		fmt.Println("有効なメールアドレス")
	}

	// 無効なメールアドレスの例
	invalidEmail := Email("notanemail")
	if !invalidEmail.IsValid() {
		fmt.Println("無効なメールアドレス（@が含まれていません）")
	}

	// Age型の使用例
	age := Age(20)
	if age.IsAdult() {
		fmt.Printf("年齢%d歳は成人です\n", age)
	}

	childAge := Age(15)
	if !childAge.IsAdult() {
		fmt.Printf("年齢%d歳は未成年です\n", childAge)
	}

	// ===== なぜこれが便利か =====
	// 1. 型に意味を持たせられる（stringではなくEmailという意図が明確）
	// 2. その型専用のメソッドを追加できる（IsValid()など）
	// 3. コンパイル時に型の混同を防げる
	// 4. コードの可読性が上がる
}
