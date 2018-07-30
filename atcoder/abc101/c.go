package main

import "fmt"

// 2018-07-30T21:26:59+0800
// 2018-07-30T21:35:20+0800
func main() {
	var N, K int
	fmt.Scanf("%d%d", &N, &K)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	x := (N - 1) / (K - 1)
	if (N-1)%(K-1) != 0 {
		x++
	}
	fmt.Println(x)
}
