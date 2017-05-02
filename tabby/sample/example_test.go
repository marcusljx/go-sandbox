package sample

import (
	"testing"

	"github.com/marcusljx/go-sandbox/tabby"
)

func TestFib(t *testing.T) {
	tabby.Func(fib).
		On(1).Expect(1).
		On(2).Expect(1).
		On(3).Expect(2).
		On(4).Expect(3).
		On(5).Expect(5).
		On(6).Expect(8).
		On(7).Expect(13).
		Run(t)
}

func TestStringReplace(t *testing.T) {
	tabby.Func(StringReplace).
		On("Hello World", "l", "++").Expect("He++++o Wor++d").
		On("Hello World", "o", "++").Expect("Hell++ W++rld").
		On("Hello World", " ", "++").Expect("Hello++World").
		Run(t)
}
