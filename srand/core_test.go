package srand

import (
	"testing"

	"github.com/marcusljx/go-sandbox/tabby"
)

func TestRandX(t *testing.T) {
	A := make(map[string]int)
	A["test"] = 1

	tabby.Func(RandomOfSet).
		Given(1).Expect(1).
		Given(false).Expect(false).
		Given([]int{1}).Expect(1).
		Given([]bool{true}).Expect(true).
		Given([]string{"hello"}).Expect("hello").
		Given([1]int{1}).Expect(1).
		Given([1]bool{true}).Expect(true).
		Given([1]string{"hello"}).Expect("hello").
		Given(A).Expect("test")
}

func BenchmarkRandomNOfSet_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNOfSet(alphaNumeric, 20)
	}
}
func BenchmarkRandomNOfSet_50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNOfSet(alphaNumeric, 50)
	}
}
func BenchmarkRandomNOfSet_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNOfSet(alphaNumeric, 100)
	}
}
func BenchmarkRandomNOfSet_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNOfSet(alphaNumeric, 1000)
	}
}
