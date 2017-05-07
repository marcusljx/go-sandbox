package slackim

import (
	"testing"

	"github.com/marcusljx/go-sandbox/tabby"
)

func TestMonospace(t *testing.T) {
	tabby.Func(Monospace).
			Given("HelloWorld").Expect("`HelloWorld`").
			Given("").Expect("``").
			Given("Hello World").Expect("`Hello World`").
			Given(123456.654321).Expect("`123456.654321`").
		Run(t)
}

func TestCode(t *testing.T) {
	tabby.Func(Code).
			Given("HelloWorld").Expect("```\nHelloWorld\n```").
			Given("").Expect("```\n\n```").
			Given("Hello World").Expect("```\nHello World\n```").
			Given(123456.654321).Expect("```\n123456.654321\n```").
		Run(t)
}

func TestHyperlink(t *testing.T) {
	tabby.Func(Hyperlink).
			Given("text", "URL").Expect("<URL|text>").
			Given("", "http://www.google.com").Expect("<http://www.google.com|>").
		Run(t)
}

func TestFindInList(t *testing.T) {
	list := []string{"Hello", "World", "Whatsup"}
	tabby.Func(findInList).
			Given(list, "hi").Expect(-1).
			Given(list, "Hello").Expect(0).
			Given(list, "World").Expect(1).
			Given(list, "Whatsup").Expect(2).
		Run(t)
}
