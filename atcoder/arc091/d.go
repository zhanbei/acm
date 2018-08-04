package main

import "fmt"

// 2018-08-04T13:47:11+0800
// 2018-08-04T14:03:18+0800
// 2018-08-04T14:40:46+0800 # @see https://arc091.contest.atcoder.jp/submissions/2940773
func main() {
	var N, K int
	fmt.Scanf("%d%d", &N, &K)
	out := 0
	for b := K + 1; b <= N; b++ {
		span := b - K
		mul := (N - K) / b
		if mul == 0 {
			out += span
		} else {
			out += span * mul
			if mul*b+b <= N {
				out += span
			} else {
				out += N + 1 - (mul*b + K)
			}
		}
	}
	if K == 0 {
		out -= N - K
	}
	fmt.Println(out)
}
