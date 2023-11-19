package calculation

// Calculator 構造体は ICalculator インターフェースを実装
type Calculator struct{}

// Add メソッド実装
func (cal *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract メソッド実装
func (cal *Calculator) Subtract(a, b int) int {
	return a - b
}
