package main

import (
	"fmt"
	"sort"
)

const MaxN = 100005
const MaxH = 1000000005

// 2018-07-29T21:25:28+0800
// 2018-07-29T21:41:10+0800
// 2018-07-29T21:57:03+0800
// 2018-07-29T22:19:12+0800
func main() {
	var N, H int
	fmt.Scanf("%d%d", &N, &H)
	var a [MaxN]int
	var b [MaxN]int
	maxa := 0
	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d", &a[i], &b[i])
		if maxa < a[i] {
			maxa = a[i]
		}
	}
	n := 0
	sort.Sort(sort.Reverse(sort.IntSlice(b[0:N])))
	for i := 0; i < N; i++ {
		if i == N-1 {
			if H-b[i] <= 0 {
				n++
				break
			}
			H -= b[i]
			n++
			if H%maxa == 0 {
				n += H / maxa
			} else {
				n += H/maxa + 1
			}
			break
		} else if b[i] > maxa {
			H -= b[i]
			n++
			if H <= 0 {
				break
			}
		} else {
			if H%maxa == 0 {
				n += H / maxa
			} else {
				n += H/maxa + 1
			}
			break
		}
	}
	fmt.Println(n)
}
