package sample

import "strings"

func fib(n int) int {
	switch n {
	case 1, 2:
		return 1
	}

	arr := make([]int, n)
	arr[0], arr[1] = 1, 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}
	return arr[n-1]
}

func StringReplace(original, a, b string) string {
	return strings.Replace(original, a, b, -1)
}
