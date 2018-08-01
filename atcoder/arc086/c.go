package main

import (
	"fmt"
	"sort"
)

// 2018-08-01T21:47:17+0800
// 2018-08-01T21:53:36+0800
// 2018-08-01T22:00:04+0800
// @see https://arc086.contest.atcoder.jp/submissions/2932146
func main() {
	var N, K int
	fmt.Scanf("%d%d", &N, &K)
	m := make(map[int]int)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		m[A[i]] ++
	}
	if len(m) <= K {
		fmt.Println(0)
		return
	}
	n := make([]int, len(m))
	i := 0
	for _, value := range m {
		n[i] = value
		i++
	}
	sort.Ints(n)

	ans := 0
	end := len(m) - K
	for i := 0; i < end; i++ {
		ans += n[i]
	}
	fmt.Println(ans)
}
