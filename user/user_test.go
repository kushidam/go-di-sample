package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockCalculator 構造体は ICalculator インターフェースのモック実装
type MockCalculator struct{}

// Add メソッド実装
func (cal *MockCalculator) Add(a, b int) int {
	return a + b
}

// Subtract メソッド実装
func (cal *MockCalculator) Subtract(a, b int) int {
	return a - b
}

func TestUserSayAdd(t *testing.T) {
	// モック作成
	mockCalculator := &MockCalculator{}

	// ユーザーを作成し、モック注入
	user := NewUser("hoge", 30, mockCalculator)

	// テスト対象のメソッドを呼び出し
	result := user.SayAdd(1, 2)

	// 期待される結果と実際の結果を比較
	assert.Equal(t, "Add: hoge 3", result)
}

func TestUserSaySubtract(t *testing.T) {
	// モック作成
	mockCalculator := &MockCalculator{}

	// ユーザーを作成し、モック注入
	user := NewUser("hoge", 30, mockCalculator)

	// テスト対象のメソッドを呼び出し
	result := user.SaySubtract(3, 2)

	// 期待される結果と実際の結果を比較
	assert.Equal(t, "Subtract: hoge 1", result)
}
