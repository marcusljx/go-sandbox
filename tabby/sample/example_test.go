package sample

import (
	"testing"

	"github.com/marcusljx/go-sandbox/tabby"
)

func TestFib(t *testing.T) {
	tabby.Func(fib).
			Given(1).Expect(1).
			Given(2).Expect(1).
			Given(3).Expect(2).
			Given(4).Expect(3).
			Given(5).Expect(5).
			Given(6).Expect(8).
			Given(7).Expect(13).
		Run(t)
}

func TestStringReplace(t *testing.T) {
	tabby.Func(StringReplace).
			Given("Hello World", "l", "++").Expect("He++++o Wor++d").
			Given("Hello World", "o", "++").Expect("Hell++ W++rld").
			Given("Hello World", " ", "++").Expect("Hello++World").
		Run(t)
}
