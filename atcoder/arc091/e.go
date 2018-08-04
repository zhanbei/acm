package main

import (
	"fmt"
	"math"
)

// 2018-08-04T14:42:31+0800
// 2018-08-04T14:54:30+0800
// 2018-08-04T15:27:19+0800 # @see https://arc091.contest.atcoder.jp/submissions/2940934
// 2018-08-04T15:32:40+0800 # @see https://arc091.contest.atcoder.jp/submissions/2940951
// 2018-08-04T15:38:28+0800 # @see https://arc091.contest.atcoder.jp/submissions/2940961
// 2018-08-04T15:40:22+0800 # @see https://arc091.contest.atcoder.jp/submissions/2940967
func main() {
	var N, A, B int
	fmt.Scanf("%d%d%d", &N, &A, &B)
	if N == 1 {
		if A != 1 || B != 1 {
			fmt.Println(-1)
			return
		} else {
			fmt.Println(1)
			return
		}
	}
	if A+B > N+1 {
		fmt.Println(-1)
		return
	}
	if false {
		min := 0
		NN := N
		root := 0
		for ; true; {
			root = int(math.Sqrt(float64(NN)))
			NN = NN - root*root
			if NN <= 2 {
				min += NN
				break
			}
			min += root
		}
		if A+B < min {
			fmt.Println(-1)
			return
		}
	} else {
		if A*B < N {
			fmt.Println(-1)
			return
		}
	}

	if B == 1 {
		for i := 0; i < N-1; i++ {
			fmt.Print(i+1, " ")
		}
		fmt.Println(N)
		return
	}
	if A == 1 {
		for i := N - 1; i > 0; i-- {
			fmt.Print(i+1, " ")
		}
		fmt.Println(1)
		return
	}

	// Now divide N to A groups while the largest one is B
	// (A-1)*span + B = N
	span := (N - B) / (A - 1)
	left := (N - B) - span*(A-2)
	out := make([]int, N)
	iout := 0
	// span <= span <= B
	for i := 0; i < A-2; i++ {
		for ii := 0; ii < span; ii++ {
			out[iout] = i*span + span - ii
			iout++
		}
	}
	i := A - 2
	for ii := 0; ii < left; ii++ {
		out[iout] = i*span + left - ii
		iout++
	}
	for ii := 0; ii < B; ii++ {
		out[iout] = i*span + left + B - ii
		iout++
	}
	for i := 0; i < N-1; i++ {
		fmt.Print(out[i], " ")
	}
	fmt.Println(out[N-1])
}
