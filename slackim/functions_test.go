package slackim

import (
	"testing"

	"github.com/marcusljx/go-sandbox/tabby"
)

func TestMonospace(t *testing.T) {
	tabby.Func(Monospace).
		On("HelloWorld").Expect("`HelloWorld`").
		On("").Expect("``").
		On("Hello World").Expect("`Hello World`").
		On(123456.654321).Expect("`123456.654321`").
		Run(t)
}

func TestCode(t *testing.T) {
	tabby.Func(Code).
		On("HelloWorld").Expect("```\nHelloWorld\n```").
		On("").Expect("```\n\n```").
		On("Hello World").Expect("```\nHello World\n```").
		On(123456.654321).Expect("```\n123456.654321\n```").
		Run(t)
}

func TestHyperlink(t *testing.T) {
	tabby.Func(Hyperlink).
		On("text", "URL").Expect("<URL|text>").
		On("", "http://www.google.com").Expect("<http://www.google.com|>").
		Run(t)
}

func TestFindInList(t *testing.T) {
	list := []string{"Hello", "World", "Whatsup"}
	tabby.Func(findInList).
		On(list, "hi").Expect(-1).
		On(list, "Hello").Expect(0).
		On(list, "World").Expect(1).
		On(list, "Whatsup").Expect(2).
		Run(t)
}
