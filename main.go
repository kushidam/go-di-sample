package main

import (
	"github.com/kushidam/go-di-sample/calculation"
	"github.com/kushidam/go-di-sample/user"
)

func main() {
	// User 構造体を生成
	user := user.NewUser("hoge", 30, &calculation.Calculator{})

	// User 構造体のメソッドを実行
	println(user.Say())
	println(user.SayAdd(1, 2))
	println(user.SaySubtract(1, 2))
}
