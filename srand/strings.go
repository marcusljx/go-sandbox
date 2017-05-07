package srand

// AlphaNumeric returns a new alphanumeric string
func AlphaNumeric(n int) string {
	s := make([]rune, n)
	for i := 0; i < n; i++ {
		s[i] = alphaNumeric[rng.Intn(n)%len(alphaNumeric)]
	}
	return string(s)
}
