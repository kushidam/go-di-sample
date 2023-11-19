package user

import "strconv"

// ICalculator インターフェースは計算機能を定義
type ICalculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
}

type User struct {
	Name   string
	SayCal ICalculator
}

func NewUser(name string, age int, sayCal ICalculator) *User {
	return &User{
		Name:   name,
		SayCal: sayCal,
	}
}

func (u *User) Say() string {
	return "Hello, " + u.Name + "!"
}

func (u *User) SayAdd(a, b int) string {
	return "Add: " + u.Name + " " + strconv.Itoa(u.SayCal.Add(a, b))
}

func (u *User) SaySubtract(a, b int) string {
	return "Subtract: " + u.Name + " " + strconv.Itoa(u.SayCal.Subtract(a, b))
}
