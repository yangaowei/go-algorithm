package main

import (
	"fmt"
)

var (
	sy = map[int]int{1: 1, 2: 5, 3: 8, 4: 9, 5: 10, 6: 17, 7: 17, 8: 20, 9: 24, 10: 30}

	X = "abcdefg"
	Y = "cdefghilmn"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func version01(n int) int {
	if n == 0 {
		return 0
	}
	q := 0
	for i := 1; i <= n; i++ {
		q = max(q, sy[i]+version01(n-i))
	}
	return q
}

func maxString(n, m int) int {
	if n < 0 || m < 0 {
		return 0
	}
	if X[n] == Y[m] {
		return maxString(n-1, m-1) + 1
	} else {
		return max(maxString(n, m-1), maxString(n-1, m))
	}
}

func main() {
	fmt.Println(version01(4))

	fmt.Println(maxString(6, 6))

	// for i := 0; i < len(X); i++ {
	// 	for j := 0; j < len(Y); j++ {
	// 		if X[i] == y[j] {
	// 			maxString(i, j) = maxString(i-1, j-1) + 1
	// 		} else {
	// 			maxString(i, j) = max(maxString(i, j-1), maxString(i-1, j))
	// 		}
	// 	}
	// }
}
