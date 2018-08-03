package main

import "fmt"

// 2018-08-03T23:14:25+0800
// 2018-08-03T23:24:36+0800
// 2018-08-03T23:54:09+0800 # @see https://arc063.contest.atcoder.jp/submissions/2939452
// 2018-08-04T00:02:23+0800 # @see https://arc063.contest.atcoder.jp/submissions/2939465
// 2018-08-04T00:23:32+0800 # @see https://arc063.contest.atcoder.jp/submissions/2939499
func main() {
	var N, T int
	fmt.Scanf("%d%d", &N, &T)
	A := make([]int, N)
	pa := make([]int, N)
	pb := make([]int, N)
	fmt.Scanf("%d", &A[0])
	pa[0] = A[0]
	for i := 1; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		if A[i] < pa[i-1] {
			pa[i] = A[i]
		} else {
			pa[i] = pa[i-1]
		}
	}
	pb[N-1] = A[N-1]
	for i := N - 2; i >= 0; i-- {
		if A[i] > pb[i+1] {
			pb[i] = A[i]
		} else {
			pb[i] = pb[i+1]
		}
	}
	t := 0
	max := -1
	// (1, 5), (2, 6)
	out := make(map[int]map[int]bool)
	for i := 1; i < N; i++ {
		t = pb[i] - pa[i-1]
		if t >= max {
			max = t
			if out[t] == nil {
				out[t] = make(map[int]bool)
			}
			out[t][pb[i]] = true
		}
	}
	fmt.Println(len(out[max]))
}

func main1() {
	var N, T int
	fmt.Scanf("%d%d", &N, &T)
	A := make([]int, N)
	AA := make([]int, N)
	iAA := 0
	fmt.Scanf("%d", &A[0])
	AA[0] = A[0]
	iAA++
	inc := true
	for i := 1; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		if i == 1 {
			inc = A[1] > A[0]
		}
		if (inc && A[i] < A[i-1]) || (!inc && A[i] > A[i-1]) {
			AA[iAA] = A[i-1]
			iAA ++
			inc = !inc
		}
	}
	AA[iAA] = A[N-1]
	iAA ++
	A = AA
	N = iAA
	pa := make([]int, N)
	pb := make([]int, N)
	pa[0] = A[0]
	for i := 1; i < N; i++ {
		if A[i] < pa[i-1] {
			pa[i] = A[i]
		} else {
			pa[i] = pa[i-1]
		}
	}
	pb[N-1] = A[N-1]
	for i := N - 2; i >= 0; i-- {
		if A[i] > pb[i+1] {
			pb[i] = A[i]
		} else {
			pb[i] = pb[i+1]
		}
	}
	out := 0
	t := 0
	max := -1
	end := -1
	for i := 1; i < N; i++ {
		t = pb[i] - pa[i-1]
		if t > max && end != pb[i] {
			max = t
			end = pb[i]
			out = 1
		} else if t == max {
			out++
		}
	}
	fmt.Println(out)
}
