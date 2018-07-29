package main

import "fmt"

var mx, my, mz int

// 2018-07-29T20:57:30+0800
// 2018-07-29T21:18:07+0800
func main() {
	var N, Y int
	fmt.Scanf("%d%d", &N, &Y)
	//dp(N, Y/1000, 0, 0, 0)
	Y = Y / 1000
	for i := 0; i <= N; i++ {
		for ii := 0; ii <= N-i; ii++ {
			z := N - i - ii
			if 10*i+5*ii+z == Y {
				fmt.Println(i, ii, z)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}

func dp(N, Y, x, y, z int) bool {
	if N == 0 {
		if Y == 0 {
			mx = x
			my = y
			mz = z
			return true
		} else {
			mx = -1
			my = -1
			mz = -1
			return false
		}
	}
	if dp(N-1, Y-10, x+1, y, z) {
		return true
	}
	if dp(N-1, Y-5, x, y+1, z) {
		return true
	}
	if dp(N-1, Y-1, x, y, z+1) {
		return true
	}
	return false
}
