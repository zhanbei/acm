package main

import "fmt"

// 2018-08-04T13:30:59+0800
// 2018-08-04T13:44:50+0800 # @see https://arc091.contest.atcoder.jp/submissions/2940605
// 2018-08-04T13:46:32+0800 # @see https://arc091.contest.atcoder.jp/submissions/2940610
func main() {
	var N, M int
	fmt.Scanf("%d%d", &N, &M)
	if N > M {
		N, M = M, N
	}
	if N == 1 {
		if M == 1 {
			fmt.Println(1)
		} else {
			fmt.Println(M - 2)
		}
	} else if N == 2 {
		fmt.Println(0)
	} else {
		fmt.Println((N - 2) * (M - 2))
	}
}
