package main

import (
	"fmt"
)

// 2018-10-25T16:28:48+0800
// 2018-10-25T16:43:57+0800 # WA @see https://arc103.contest.atcoder.jp/submissions/3466776
// 2018-10-25T16:45:49+0800 # WA @see https://arc103.contest.atcoder.jp/submissions/3466791
// 2018-10-25T17:03:41+0800 # Failed to pause on fmt#Scanf() when debugging.
// 2018-10-25T17:07:57+0800 # @see https://arc103.contest.atcoder.jp/submissions/3466886
func main() {
	var n int
	fmt.Scanf("%d", &n)
	//v := make([]int, n)
	a := make([]int, n/2)
	b := make([]int, n/2)
	ma := make(map[int]int, n/2)
	aa1 := -1
	aa2 := -1
	mb := make(map[int]int, n/2)
	bb1 := -1
	bb2 := -1
	for i := 0; i < n/2; i++ {
		fmt.Scanf("%d", &a[i])
		ma[a[i]]++
		if aa1 != a[i] && ma[a[i]] >= ma[aa1] {
			aa2 = aa1
			aa1 = a[i]
		}
		fmt.Scanf("%d", &b[i])
		mb[b[i]]++
		if bb1 != b[i] && mb[b[i]] >= mb[bb1] {
			bb2 = bb1
			bb1 = b[i]
		}
	}
	if aa1 != bb1 {
		fmt.Println(n - ma[aa1] - mb[bb1])
		return
	}
	O1 := n - ma[aa1] - mb[bb2]
	O2 := n - ma[aa2] - mb[bb1]
	if O1 > O2 {
		fmt.Println(O2)
	} else {
		fmt.Println(O1)
	}
}
