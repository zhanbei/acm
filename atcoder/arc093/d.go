package main

import "fmt"

// 2018-08-09T23:16:13+0800
// 2018-08-09T23:34:23+0800
// 2018-08-10T00:10:30+0800 # @see https://arc093.contest.atcoder.jp/submissions/2977167
func main() {
	var A, B int
	fmt.Scanf("%d%d", &A, &B)
	sq := make([][]bool, 100)
	h, w := 100, 100
	for i := 0; i < h; i++ {
		sq[i] = make([]bool, w)
	}

	A -= 1
	B -= 1

	for h := 0; A > 0; A -= w / 2 {
		for ii := 0; ii < w; ii += 2 {
			if ii/2 >= A {
				break
			}
			sq[h][ii] = true
		}
		h += 2
	}
	for h := 51; B > 0; B -= w / 2 {
		for ii := 0; ii < w; ii += 2 {
			if ii/2 >= B {
				break
			}
			sq[h][ii] = true
		}
		h += 2
	}

	fmt.Println(h, w)
	for i := 0; i < h; i++ {
		for ii := 0; ii < w; ii++ {
			if i < 50 {
				if sq[i][ii] {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			} else {
				if sq[i][ii] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}
