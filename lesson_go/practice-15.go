package main

import (
	"log"
	"os"
)

func lessonDefer() {
	// deferは関数の実行が終わった際に実行される
	f, err := os.Open("no_exist.txt")
	if err != nil {
		log.Fatal(err)
	}
	// これによってこの関数が存続するまで、ファイルのリソースは開いたままになる。
	defer f.Close()

}
