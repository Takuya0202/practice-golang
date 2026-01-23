package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func lessonPointer() {
	// ポインタは&を使って参照し、*で実態を参照できる
	v := 1
	p := &v
	*p = 2
	fmt.Println(v)

	// personのポインタ作成
	var p2 *Person

	p2 = &Person{
		Name: "山田",
		Age:  20,
	}
	fmt.Printf("変数pに格納されているアドレス :%p", p2) // メモリが表示される
	fmt.Println()

	// newで作成したstructはポインタで返す。
	person := new(Person)
	person.Name = "佐藤"
	person.Age = 30
	fmt.Println(person)

	// ポインタのメソッド

}
