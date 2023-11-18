package main

import (
	"github.com/kushidam/go-di-sample/calculation"
)

func main() {
	println("Add", calculation.Add(1,2))
	println("Subtract", calculation.Subtract(5,2))
}
