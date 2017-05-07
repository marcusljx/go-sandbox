package srand

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func isInSet(set []rune, r rune) bool {
	for _, s := range set {
		if s == r {
			return true
		}
	}
	return false
}

func TestAlphaNumeric(t *testing.T) {
	var s string
	f := func() bool {
		for _, c := range s {
			if !isInSet(alphaNumeric, c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < 200000; i++ {
		s = AlphaNumeric(20)
		require.Condition(t, f, "s was [%v]", s)
	}
}

func BenchmarkAlphaNumeric_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaNumeric(20)
	}
}
func BenchmarkAlphaNumeric_50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaNumeric(50)
	}
}
func BenchmarkAlphaNumeric_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaNumeric(100)
	}
}
func BenchmarkAlphaNumeric_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaNumeric(1000)
	}
}
