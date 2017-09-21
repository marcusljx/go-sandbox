package xjlm

import "log"

// CheckErrPanic panics if the given error is non-nil
func CheckErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckErrLog logs the error if the given error is non-nil
func CheckErrLog(err error) {
	if err != nil {
		log.Print(err)
	}
}

// CheckErrPanic ignores the given error.
// It is equivalent to a noop.
func CheckErrIgnore(err error) {}
