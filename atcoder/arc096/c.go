package main

import "fmt"

// 2018-08-01T15:35:44+0800
// 2018-08-01T15:55:34+0800
// 2018-08-01T16:14:00+0800
// https://arc096.contest.atcoder.jp/submissions/2930956
func main() {
	var A, B, C, X, Y int
	fmt.Scanf("%d%d%d%d%d", &A, &B, &C, &X, &Y)
	if 2*C >= A+B {
		n := X*A + Y*B
		fmt.Println(n)
	} else {
		if X > Y {
			A, B = B, A
			X, Y = Y, X
		}
		// A <= B
		// z = 2*X, Y -= X
		n := C * 2 * X
		Y -= X
		if 2*C >= B {
			n += Y * B
		} else {
			n += C * 2 * Y
		}
		fmt.Println(n)
	}
}
