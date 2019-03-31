package main

import (
	"fmt"
	"sort"
)

// 2019-03-31T23:37:06+0800
// 2019-03-31T23:44:53+0800
// 2019-03-31T23:56:49+0800 # AC @see https://atcoder.jp/contests/arc079/submissions/4790466
func main() {
	var N int
	fmt.Scanf("%d", &N)
	a := make([]int, N)
	n := 0
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
		n += a[i]
	}
	out := solve(a)
	fmt.Println(out)
}

func solve(a []int) int {
	N := len(a)
	sort.Ints(a)
	if a[N-1] < N {
		return 0
	}
	total := 0
	for i := N - 1; i >= 0; i-- {
		if a[i] < N {
			break
		}
		times := a[i] / N
		left := a[i] % N
		total += times

		for ii := 0; ii < N; ii++ {
			a[ii] += times
		}
		a[i] = left
	}
	return total + solve(a)
}
