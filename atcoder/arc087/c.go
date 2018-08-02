package main

import (
	"fmt"
)

// 2018-08-02T12:25:38+0800
// 2018-08-02T12:28:40+0800
// 2018-08-02T12:32:30+0800
// @see https://arc087.contest.atcoder.jp/submissions/2933536
func main() {
	var N int
	fmt.Scanf("%d", &N)
	a := make([]int, N)
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
		m[a[i]] ++
	}
	out := 0
	for key, value := range m {
		if value > key {
			out += value - key
		} else if value < key {
			out += value
		}
	}
	fmt.Println(out)
}
