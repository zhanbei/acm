package main

import (
	"fmt"
	"sort"
)

var N int
var xs []int
var ys []int
var sortedX []int
var sortedY []int
// x, y, -> index
var xys = make(map[int]map[int]int)
var yxs = make(map[int]map[int]int)
// Whether points being checked.
var flags = make(map[int]map[int]bool)
var checkedPoints = make(map[int]bool)
// Reachable points.
//var points map[int]bool
var pace int
var Out int

// 2018-11-04T01:06:38+0800
// 2018-11-04T01:15:59+0800
// 2018-11-04T01:51:52+0800 # TLE @see https://arc065.contest.atcoder.jp/submissions/3527473
// 2018-11-04T02:04:30+0800 Try limiting x and y.
// 2018-11-04T03:08:06+0800 # TLE but much better :) @see https://arc065.contest.atcoder.jp/submissions/3527581
func main() {
	var a, b int
	fmt.Scanf("%d%d%d", &N, &a, &b)
	a--
	b--
	flags = make(map[int]map[int]bool, N)
	//points = make(map[int]bool, N)
	xs = make([]int, N)
	ys = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d", &xs[i], &ys[i])
		if xys[xs[i]] == nil {
			xys[xs[i]] = make(map[int]int)
		}
		xys[xs[i]][ys[i]] = i
		if yxs[ys[i]] == nil {
			yxs[ys[i]] = make(map[int]int)
		}
		yxs[ys[i]][xs[i]] = i
		flags[i] = make(map[int]bool)
	}

	sortedX = make([]int, len(xys))
	i := 0
	for k := range xys {
		sortedX[i] = k
		i++
	}
	i = 0
	sortedY = make([]int, len(yxs))
	for k := range yxs {
		sortedY[i] = k
		i++
	}
	sort.Ints(sortedX)
	sort.Ints(sortedY)

	pace = length(a, b)
	//points[a] = true
	//points[b] = true
	flags[a] = make(map[int]bool)
	flags[b] = make(map[int]bool)
	flags[a][b] = true
	flags[b][a] = true
	//Out = 2
	findx(a)
	findx(b)

	fmt.Println(Out / 2)
}

func findx(cur int) {
	if checkedPoints[cur] {
		return
	}
	checkedPoints[cur] = true
	start := sort.SearchInts(sortedX, xs[cur]-pace)
	//fmt.Println("Checking", cur, "starting from:", start, "of", sortedX)
	for i := start; i < len(sortedX) && sortedX[i] <= xs[cur]+pace; i++ {
		offset := pace - abs(sortedX[i]-xs[cur])
		if offset == 0 {
			//fmt.Println(cur, "expects", sortedX[i], "-", ys[cur]+offset, ys[cur]-offset)
			if next, ok := xys[sortedX[i]][ys[cur]]; ok {
				//fmt.Println("found", cur, "with", next, ":", sortedX[i], "-", ys[cur], offset)
				Out++
				findx(next)
			}
			continue
		}
		//fmt.Println(cur, "expects", sortedX[i], "-", ys[cur]+offset, ys[cur]-offset)
		if next, ok := xys[sortedX[i]][ys[cur]+offset]; ok {
			//fmt.Println("found", cur, "with", next, ":", sortedX[i], "-", ys[cur]+offset, offset)
			Out++
			findx(next)
		}
		if next, ok := xys[sortedX[i]][ys[cur]-offset]; ok {
			//fmt.Println("found", cur, "with", next, ":", sortedX[i], "-", ys[cur]-offset, offset)
			Out++
			findx(next)
		}
	}
}

func main2() {
	var a, b int
	fmt.Scanf("%d%d%d", &N, &a, &b)
	a--
	b--
	flags = make(map[int]map[int]bool, N)
	//points = make(map[int]bool, N)
	xs = make([]int, N)
	ys = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d", &xs[i], &ys[i])
		flags[i] = make(map[int]bool)
	}

	pace = length(a, b)
	//points[a] = true
	//points[b] = true
	flags[a] = make(map[int]bool)
	flags[b] = make(map[int]bool)
	flags[a][b] = true
	flags[b][a] = true
	Out = 1
	find(a)
	find(b)

	fmt.Println(Out)
}

func find(cur int) {
	//if flags[cur] == nil {
	//	flags[cur] = make(map[int]bool)
	//}
	for i := 0; i < N; i++ {
		if flags[cur][i] {
			// Already checked.
			continue
		}
		flags[i][cur] = true
		flags[cur][i] = true
		if length(cur, i) == pace {
			//fmt.Println("cur, i:", cur, i)
			Out++
			//points[i] = true
			find(i)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func length(a, b int) int {
	return abs(xs[a]-xs[b]) + abs(ys[a]-ys[b])
}
