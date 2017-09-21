package xjlm

import (
	"path/filepath"
	"runtime"
)

// CurrentDir returns the current directory of the calling code
func CurrentDir() string {
	_, f, _, _ := runtime.Caller(1)
	return filepath.Dir(f)
}

// RelativeDir returns the fully qualified path of the given relative directory path
func RelativeDir(relativePath string) string {
	_, f, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(f), relativePath)
}
