package tabby

import "fmt"

const (
	_logSpacingPattern = "\t\t\t%8v : %v"
)

func formatKeyValueOutput(key, value interface{}) string {
	return fmt.Sprintf(_logSpacingPattern, key, value)
}
