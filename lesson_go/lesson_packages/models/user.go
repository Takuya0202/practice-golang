// modelsパッケージ - データモデルを定義するパッケージ
package models

import "fmt"

// User構造体 - 公開（大文字で始まる）
// 他のパッケージから使える
type User struct {
	// 公開フィールド（大文字で始まる）
	// 他のパッケージから直接アクセスできる
	ID    int
	Name  string
	Email string

	// 非公開フィールド（小文字で始まる）
	// このパッケージ内でのみアクセス可能
	age int
}

// NewUser - 公開関数（大文字で始まる）
// User構造体を作成するコンストラクタ的な関数
// 他のパッケージから呼び出せる
func NewUser(id int, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		age:   0, // 非公開フィールドは初期値
	}
}

// GetAge - 公開メソッド（大文字で始まる）
// 非公開フィールドageにアクセスするための公開メソッド
// 他のパッケージから呼び出せる
func (u *User) GetAge() int {
	return u.age
}

// SetAge - 公開メソッド
// 非公開フィールドageを設定するための公開メソッド
// バリデーションも含めることができる
func (u *User) SetAge(age int) {
	// 負の値は0に設定（バリデーション例）
	if age < 0 {
		age = 0
	}
	u.age = age
}

// GetInfo - 公開メソッド
// ユーザー情報を文字列で返す
func (u *User) GetInfo() string {
	return fmt.Sprintf("ID: %d, 名前: %s, メール: %s, 年齢: %d", u.ID, u.Name, u.Email, u.age)
}

// validateEmail - 非公開関数（小文字で始まる）
// このパッケージ内でのみ使える
// 他のパッケージからは呼び出せない
func validateEmail(email string) bool {
	return len(email) > 0 && contains(email, "@")
}

// contains - 非公開関数
// このパッケージ内でのみ使えるヘルパー関数
func contains(s, substr string) bool {
	// 簡易実装（実際はstrings.Containsを使う）
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// ValidateUserEmail - 公開関数
// 非公開関数validateEmailを呼び出す公開関数
// 他のパッケージからも使える
func ValidateUserEmail(email string) bool {
	return validateEmail(email)
}
