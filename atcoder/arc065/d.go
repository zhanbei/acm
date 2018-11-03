package main

import "fmt"

type City struct {
	Id int
	// Connected cities by roads
	ByRoads    map[int]bool
	allByRoads map[int]bool

	ByRailways    map[int]bool
	allByRailways map[int]bool

	sum int
}

var m map[int]bool
var cities []*City

//func main() {
//	N := 10000
//	fmt.Println(N, N/2, N/2)
//	for i := 0; i < N/2; i++ {
//		fmt.Println(i+1, i+2)
//	}
//}

// 2018-11-03T18:05:12+0800
// 2018-11-03T18:27:41+0800
// 2018-11-03T18:57:00+0800 # TLE @see https://arc065.contest.atcoder.jp/submissions/3525055
// 2018-11-03T19:31:12+0800 # TLE @see https://arc065.contest.atcoder.jp/submissions/3525205
// 2018-11-03T19:39:21+0800 # TLE @see https://arc065.contest.atcoder.jp/submissions/3525251
// 2018-11-03T19:54:40+0800 Giving up.
func main() {
	var N, K, L int
	fmt.Scanf("%d%d%d", &N, &K, &L)
	cities = make([]*City, N)
	for i := 0; i < N; i++ {
		cities[i] = &City{
			i,
			make(map[int]bool), nil,
			make(map[int]bool), nil,
			0,
		}
	}
	var a, b int
	for i := 0; i < K; i++ {
		fmt.Scanf("%d%d", &a, &b)
		a--
		b--
		cities[a].ByRoads[b] = true
		cities[b].ByRoads[a] = true
	}
	for i := 0; i < L; i++ {
		fmt.Scanf("%d%d", &a, &b)
		a--
		b--
		cities[a].ByRailways[b] = true
		cities[b].ByRailways[a] = true
	}
	//p := make([]int, K)
	//q := make([]int, K)
	//r := make([]int, L)
	//s := make([]int, L)
	//for i := 0; i < K; i++ {
	//	fmt.Scanf("%d%d", &p[i], &q[i])
	//	p[i]--
	//	q[i]--
	//	cities[p[i]].ByRoads[q[i]] = true
	//	cities[q[i]].ByRoads[p[i]] = true
	//}
	//for i := 0; i < L; i++ {
	//	fmt.Scanf("%d%d", &r[i], &s[i])
	//	r[i]--
	//	s[i]--
	//	cities[r[i]].ByRailways[s[i]] = true
	//	cities[s[i]].ByRailways[r[i]] = true
	//}

	for i := 0; i < N; i++ {
		if cities[i].allByRoads != nil {
			continue
		}
		// Create a new group of connected cities.
		m = make(map[int]bool)
		roads(i)
	}

	for i := 0; i < N; i++ {
		if cities[i].allByRailways != nil {
			continue
		}
		// Create a new group of connected cities.
		m = make(map[int]bool)
		rails(i)
	}

	for i := 0; i < N; i++ {
		city := cities[i]
		for k := range city.allByRoads {
			if city.allByRailways[k] {
				city.sum++
			}
		}
	}

	for i := 0; i < N-1; i++ {
		fmt.Print(cities[i].sum, " ")
	}
	fmt.Println(cities[N-1].sum)
}

func roads(i int) {
	//fmt.Println("road:", i, len(m))
	//if cities[i].allByRoads != nil {
	//	panic("now")
	//}
	m[i] = true
	cities[i].allByRoads = m
	for ii := range cities[i].ByRoads {
		if !m[ii] {
			roads(ii)
		}
	}
}

func rails(i int) {
	//fmt.Println("rail:", i, len(cities[i].allByRailways))
	//if cities[i].allByRailways != nil {
	//	panic("now")
	//}
	m[i] = true
	cities[i].allByRailways = m
	for ii := range cities[i].ByRailways {
		if !m[ii] {
			rails(ii)
		}
	}
}
