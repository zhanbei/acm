package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var A, B int
	fmt.Scanf("%d%d", &A, &B)
	fmt.Printf("%d", max(A*B, max(A+B, A-B)))
}
