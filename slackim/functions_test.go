package slackim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMonospace(t *testing.T) {
	require.Equal(t, "`HelloWorld`", Monospace("HelloWorld"))
	require.Equal(t, "``", Monospace(""))
	require.Equal(t, "`Hello World`", Monospace("Hello World"))
	require.Equal(t, "`123456.654321`", Monospace(123456.654321))
}

func TestCode(t *testing.T) {
	require.Equal(t, "```\nHelloWorld\n```", Code("HelloWorld"))
	require.Equal(t, "```\n\n```", Code(""))
	require.Equal(t, "```\nHello World\n```", Code("Hello World"))
	require.Equal(t, "```\n123456.654321\n```", Code(123456.654321))
}

func TestHyperlink(t *testing.T) {
	require.Equal(t, "<URL|text>", Hyperlink("text", "URL"))
	require.Equal(t, "<http://www.google.com|>", Hyperlink("", "http://www.google.com"))
}

func TestFindInList(t *testing.T) {
	list := []string{"Hello", "World", "Whatsup"}
	require.Equal(t, -1, findInList(list, "hi"))
	require.Equal(t, 0, findInList(list, "Hello"))
	require.Equal(t, 1, findInList(list, "World"))
	require.Equal(t, 2, findInList(list, "Whatsup"))
}
