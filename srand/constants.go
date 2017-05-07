package srand

var (
	numeric       = []rune("0123456789")
	alphabetLower = []rune("abcdefghijklmnopqrstuvwxyz")
	alphabetUpper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphabet      = append(alphabetLower, alphabetUpper...)
	alphaNumeric  = append(alphabet, numeric...)
)
