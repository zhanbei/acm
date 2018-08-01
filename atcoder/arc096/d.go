package main

import "fmt"

// 2018-08-01T16:32:46+0800
// 2018-08-01T16:45:39+0800
// 2018-08-01T16:57:24+0800
// Giving up and looking up for help
// 2018-08-01T17:19:01+0800
// 2018-08-01T17:36:10+0800
// 2018-08-01T18:30:08+0800
// @see https://arc096.contest.atcoder.jp/submissions/2931406
func main() {
	var N, C int
	fmt.Scanf("%d%d", &N, &C)
	x := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d", &x[i], &v[i])
	}
	// sum, max
	ma := make([]int, N)
	mb := make([]int, N)
	t := 0
	sa := v[0]
	ma[0] = v[0] - x[0]
	sb := v[N-1]
	mb[N-1] = v[N-1] - C + x[N-1]
	for i, ii := 1, N-2; i < N; i++ {
		sa += v[i]
		t = sa - x[i]
		if t > ma[i-1] {
			ma[i] = t
		} else {
			ma[i] = ma[i-1]
		}
		sb += v[ii]
		t = sb - C + x[ii]
		if t > mb[ii+1] {
			mb[ii] = t
		} else {
			mb[ii] = mb[ii+1]
		}
		ii--
	}
	t, n := 0, 0
	back := 0
	for i := -1; i < N; i++ {
		if i == -1 {
			t = 0
			back = 0
		} else {
			t = ma[i]
			back = x[i]
		}
		if i+1 < N && mb[i+1]-back > 0 {
			t += mb[i+1] - back
		}
		if t > n {
			n = t
		}
	}
	for i := N; i >= 0; i-- {
		if i == N {
			t = 0
			back = 0
		} else {
			t = mb[i]
			back = C - x[i]
		}
		if i > 0 && ma[i-1]-back > 0 {
			t += ma[i-1] - back
		}
		if t > n {
			n = t
		}
	}
	fmt.Println(n)
}

func wrongSolution() {
	var N, C int
	fmt.Scanf("%d%d", &N, &C)
	x := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d", &x[i], &v[i])
	}
	var sa, na, ia, sb, nb, ib int
	t := 0
	for i, ii := 0, N-1; i < N; i++ {
		sa += v[i]
		t = sa - x[i]
		if t > na {
			na = t
			ia = i
		}
		sb += v[ii]
		t = sb - C + x[ii]
		if t > nb {
			nb = t
			ib = ii
		}
		ii--
	}
	n := 0
	if ia >= ib {
		fmt.Println("ia >= ib", ia, ib)
		if na > nb {
			n = na
		} else {
			n = nb
		}
	} else {
		if na > nb {
			n = na
			if nb-x[ia] > 0 {
				n += nb - x[ia]
			}
		} else {
			n = nb
			if na-C+x[ib] > 0 {
				n += nb - C + x[ib]
			}
		}
	}
	fmt.Println(n)
}
