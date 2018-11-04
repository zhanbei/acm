package main

import (
	"fmt"
	"sort"
)

// 2018-11-04T14:37:39+0800 # TRY Caching the distance.
// 2018-11-04T16:42:44+0800 # TLE but passed 3 more cases. @see https://arc065.contest.atcoder.jp/submissions/3529702
// 2018-11-04T17:05:16+0800 # TLE @see https://arc065.contest.atcoder.jp/submissions/3529825
func main() {
	var N, a, b int
	fmt.Scanf("%d%d%d", &N, &a, &b)
	a--
	b--
	//points = make(map[int]bool, N)
	xs := make([]int, N)
	ys := make([]int, N)
	// dis -> xs
	dises0 := make(map[int][]int, N)
	// dis -> xs -> index
	indexesByDis0 := make(map[int]map[int]int, N)
	// whether is sorted
	flags0 := make(map[int]bool, N)
	// dis = x -y -> xs
	dises2 := make(map[int][]int, N)
	indexesByDis2 := make(map[int]map[int]int, N)
	flags2 := make(map[int]bool, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d", &xs[i], &ys[i])
		dis := xs[i] + ys[i]
		if dises0[dis] == nil {
			dises0[dis] = make([]int, 0)
			indexesByDis0[dis] = make(map[int]int)
		}
		dises0[dis] = append(dises0[dis], xs[i])
		indexesByDis0[dis][xs[i]] = i

		dis = xs[i] - ys[i]
		if dises2[dis] == nil {
			dises2[dis] = make([]int, 0)
			indexesByDis2[dis] = make(map[int]int)
		}
		dises2[dis] = append(dises2[dis], xs[i])
		indexesByDis2[dis][xs[i]] = i
	}

	pace := abs(xs[a]-xs[b]) + abs(ys[a]-ys[b])
	checkedPoints := make(map[int]bool)
	Out := 0

	var cal func(cur, xFrom int)
	cal = func(cur, xFrom int) {
		//fmt.Println("found", cur+1, "from", xFrom+1)
		if checkedPoints[cur] {
			return
		}
		checkedPoints[cur] = true
		dist := xs[cur] + ys[cur] - pace
		if len(dises0[dist]) > 0 {
			if !flags0[dist] {
				sort.Ints(dises0[dist])
				flags0[dist] = true
			}
			start := sort.SearchInts(dises0[dist], xs[cur]-pace)
			//fmt.Println("1 Looking in", dises0[dist], start)
			if start < len(dises0[dist]) && dises0[dist][start] == xs[cur]-pace {
				start++
			}
			for i := start; i < len(dises0[dist]) && dises0[dist][i] < xs[cur]; i++ {
				Out++
				cal(indexesByDis0[dist][dises0[dist][i]], cur)
			}
		}

		dist = xs[cur] + ys[cur] + pace
		if len(dises0[dist]) > 0 {
			if !flags0[dist] {
				sort.Ints(dises0[dist])
				flags0[dist] = true
			}
			start := sort.SearchInts(dises0[dist], xs[cur])
			//fmt.Println("2 Looking in", dises0[dist], start)
			if start < len(dises0[dist]) && dises0[dist][start] == xs[cur] {
				start++
			}
			for i := start; i < len(dises0[dist]) && dises0[dist][i] < xs[cur]+pace; i++ {
				Out++
				cal(indexesByDis0[dist][dises0[dist][i]], cur)
			}
		}

		dist = xs[cur] - ys[cur] - pace
		if len(dises2[dist]) > 0 {
			if !flags2[dist] {
				sort.Ints(dises2[dist])
				flags2[dist] = true
			}
			start := sort.SearchInts(dises2[dist], xs[cur]-pace)
			//fmt.Println("3 Looking in", dises2[dist], start)
			for i := start; i < len(dises2[dist]) && dises2[dist][i] <= xs[cur]; i++ {
				Out++
				cal(indexesByDis2[dist][dises2[dist][i]], cur)
			}
		}

		dist = xs[cur] - ys[cur] + pace
		if len(dises2[dist]) > 0 {
			if !flags2[dist] {
				sort.Ints(dises2[dist])
				flags2[dist] = true
			}
			start := sort.SearchInts(dises2[dist], xs[cur])
			//fmt.Println("4 Looking in", dises2[dist], start)
			for i := start; i < len(dises2[dist]) && dises2[dist][i] <= xs[cur]+pace; i++ {
				Out++
				cal(indexesByDis2[dist][dises2[dist][i]], cur)
			}
		}
	}

	cal(a, -2)
	//cal(b, 0)

	fmt.Println(Out / 2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
