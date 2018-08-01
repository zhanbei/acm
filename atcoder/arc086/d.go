package main

import (
	"fmt"
	"math"
)

// 2018-08-01T22:01:05+0800
// 2018-08-01T22:11:37+0800
// Giving up and looking for help
// 2018-08-01T22:13:57+0800
// 2018-08-01T22:17:29+0800
// 2018-08-01T22:35:08+0800
// @see https://arc086.contest.atcoder.jp/submissions/2932250
func main() {
	var N int
	fmt.Scanf("%d", &N)
	a := make([]int, N)
	max := math.MinInt32
	maxi := 0
	min := math.MaxInt32
	mini := 0
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
		if a[i] > max {
			max = a[i]
			maxi = i
		}
		if a[i] < min {
			min = a[i]
			mini = i
		}
	}
	if min > 0 {
		fmt.Println(N - 1)
		for i := 0; i < N-1; i++ {
			fmt.Println(i+1, i+2)
		}
	} else if max < 0 {
		fmt.Println(N - 1)
		for i := N - 2; i >= 0; i-- {
			fmt.Println(i+2, i+1)
		}
	} else {
		if max > -min {
			// Set all to non-negative numbers.
			fmt.Println(2*N - 1)
			for i := 0; i < N; i++ {
				fmt.Println(maxi+1, i+1)
			}
			for i := 0; i < N-1; i++ {
				fmt.Println(i+1, i+2)
			}
		} else {
			fmt.Println(2*N - 1)
			for i := 0; i < N; i++ {
				fmt.Println(mini+1, i+1)
			}
			for i := N - 2; i >= 0; i-- {
				fmt.Println(i+2, i+1)
			}
		}
	}
}
