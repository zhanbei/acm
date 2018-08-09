package main

import "fmt"

// 2018-08-09T23:03:56+0800
// 2018-08-09T23:08:44+0800
// 2018-08-09T23:15:20+0800 # @see https://arc093.contest.atcoder.jp/submissions/2977020
func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N+2)
	sum := 0
	sum += diff(A[0], A[1])
	for i := 1; i < N+1; i++ {
		fmt.Scanf("%d", &A[i])
		sum += diff(A[i], A[i-1])
	}
	sum += diff(A[N], A[N+1])
	for i := 1; i < N+1; i++ {
		fmt.Println(sum + diff(A[i-1], A[i+1]) - diff(A[i-1], A[i]) - diff(A[i], A[i+1]))
	}
}

func diff(a, b int) int {
	c := a - b
	if c < 0 {
		return -c
	}
	return c
}
