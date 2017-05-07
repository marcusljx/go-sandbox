package srand

import (
	"math/rand"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func randomOfRuneSet(in []rune) rune {
	return in[rng.Intn(len(in))]
}

func randomNOfSet(in []rune, n int) (result []rune) {
	result = make([]rune, n)
	for i := 0; i < n; i++ {
		result[i] = randomOfRuneSet(in)
	}
	return
}
