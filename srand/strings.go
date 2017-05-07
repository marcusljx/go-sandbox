package srand

// AlphaNumeric returns an alphanumeric string
func AlphaNumeric(n int) string {
	return string(randomNOfSet(alphaNumeric, n))
}
