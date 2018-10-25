package main

import (
	"fmt"
	"strings"
)

// 2018-10-25T17:11:07+0800
// 2018-10-25T17:34:04+0800 # Starting solving the small case.
// 2018-10-25T17:50:33+0800 # WA for big case @see https://arc103.contest.atcoder.jp/submissions/3467176
func main() {
	solveSmall()
}

func solveSmall() {
	var N int
	fmt.Scanf("%d", &N)
	X := make([]int, N)
	Y := make([]int, N)
	fmt.Scanf("%d%d", &X[0], &Y[0])
	isEven := (X[0]+Y[0])%2 == 0
	noAnswer := false
	for i := 1; i < N; i++ {
		fmt.Scanf("%d%d", &X[i], &Y[i])
		if noAnswer {
			continue
		}
		if (X[i]+Y[i])%2 == 0 != isEven {
			noAnswer = true
		}
		if abs(X[i]) > 10 || abs(Y[i]) > 10 {
			noAnswer = true
		}
	}
	if noAnswer {
		fmt.Println(-1)
		return
	}
	m := 20
	if !isEven {
		m = 19
	}
	fmt.Println(m)
	for i := 0; i < m-1; i++ {
		fmt.Print(1, " ")
	}
	fmt.Println(1)

	for i := 0; i < N; i++ {
		extra := m - abs(X[i]) - abs(Y[i])
		fmt.Print(strings.Repeat("LR", extra/2))
		if X[i] > 0 {
			fmt.Print(strings.Repeat("R", X[i]))
		} else {
			fmt.Print(strings.Repeat("L", -X[i]))
		}
		if Y[i] > 0 {
			fmt.Print(strings.Repeat("U", Y[i]))
		} else {
			fmt.Print(strings.Repeat("D", -Y[i]))
		}
		fmt.Println()
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
