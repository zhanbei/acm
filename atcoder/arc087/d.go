package main

import (
	"fmt"
	"strings"
)

const Right = 1
const Up = 2
const Left = 3
const Down = 4

var F = "F"[0]
var T = "T"[0]
var tx, ty int

// 2018-08-02T12:33:59+0800
// 2018-08-02T12:43:06+0800
// 2018-08-02T13:01:55+0800
// dfs1() -> @see https://arc087.contest.atcoder.jp/submissions/2933611
// 2018-08-02T13:32:53+0800
// dfs() -> @see https://arc087.contest.atcoder.jp/submissions/2933712
// dfs() -> @see https://arc087.contest.atcoder.jp/submissions/2933766
// 2018-08-02T13:57:26+0800 # Giving up and looking for help.
// 2018-08-02T14:04:20+0800
// 2018-08-02T15:14:25+0800 # @see https://arc087.contest.atcoder.jp/submissions/2934062
// 2018-08-02T15:35:00+0800 # @see https://arc087.contest.atcoder.jp/submissions/2934106
// 2018-08-02T21:40:34+0800 # Looked for help and start coding.
// 2018-08-02T21:54:53+0800 # AC Using dp with cache.
func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d%d", &tx, &ty)
	s = "F" + s + "TF"
	// how many continuous f, how many continuous t, ...
	// Length of vh will be ii+1.
	vh := make([]int, len(s)+1)
	// Length of dx will be idx.
	dx := make([]int, len(s)+1)
	// Length of dy will be idy.
	dy := make([]int, len(s)+1)
	isx := true
	idx, idy := 0, 0
	ii := 0
	vh[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if s[i] == F {
				if isx {
					dx[idx] = vh[ii-1]
					idx ++
				} else {
					dy[idy] = vh[ii-1]
					idy ++
				}
				if vh[ii]%2 != 0 {
					isx = !isx
				}
			}
			ii++
		}
		vh[ii]++
	}
	x, y := vh[0], 0
	// x -> tx
	// y -> ty
	initx := make(map[int]bool, 1)
	initx[x] = true
	inity := make(map[int]bool, 1)
	inity[y] = true
	if dpForMap(dx[1:idx], initx)[tx] && dpForMap(dy[0:idy], inity)[ty] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dpForMap(dd []int, ends map[int]bool) map[int]bool {
	if len(dd) == 0 {
		return ends
	}
	m := make(map[int]bool, len(ends)*2)
	for k := range ends {
		m[k-dd[0]] = true
		m[k+dd[0]] = true
	}
	return dpForMap(dd[1:], m)
}

func dpCacheMain() {
	var dx, dy []int
	var idx, x, idy, y int
	if dpx(dx[1:idx], x) && dpy(dy[0:idy], y) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var cacheDx = make(map[int]int)

func dpx(dx []int, x int) bool {
	if len(dx) == 0 {
		return x == tx
	}
	l := len(dx[1:]) * 10000
	a, b := l+x+dx[0], l+x-dx[0]
	t := false
	switch cacheDx[a] {
	case 1:
		t = true
	case -1:
		t = false
	default:
		t = dpx(dx[1:], x+dx[0])
		if t {
			cacheDx[a] = 1
		} else {
			cacheDx[a] = -1
		}
	}
	if t {
		return true
	}
	switch cacheDx[b] {
	case 1:
		return true
	case -1:
		return false
	default:
		t = dpx(dx[1:], x-dx[0])
		if t {
			cacheDx[a] = 1
		} else {
			cacheDx[a] = -1
		}
		return t
	}
}

var cacheDy = make(map[int]int)

func dpy(dy []int, y int) bool {
	if len(dy) == 0 {
		return y == ty
	}
	l := len(dy[1:]) * 10000
	a, b := l+y+dy[0], l+y-dy[0]
	t := false
	switch cacheDy[a] {
	case 1:
		t = true
	case -1:
		t = false
	default:
		t = dpy(dy[1:], y+dy[0])
		if t {
			cacheDy[a] = 1
		} else {
			cacheDy[a] = -1
		}
	}
	if t {
		return true
	}
	switch cacheDy[b] {
	case 1:
		return true
	case -1:
		return false
	default:
		t = dpy(dy[1:], y-dy[0])
		if t {
			cacheDy[a] = 1
		} else {
			cacheDy[a] = -1
		}
		return t
	}
}

var prefixes []int
var slength int

func main1() {
	var s, news string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d%d", &tx, &ty)
	// Replace 4T to 2T.
	for ; true; {
		news = strings.Replace(s, "TTT", "T", -1)
		if len(news) == len(s) {
			break
		}
		s = news
	}
	slength = len(s)
	// Calculate the prefixes of 'F'.
	prefixes = make([]int, len(s))
	if s[len(s)-1] == F {
		prefixes[len(s)-1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == F {
			prefixes[i] = prefixes[i+1] + 1
		} else {
			prefixes[i] = prefixes[i+1]
		}
	}
	if dfs(s, Right, 0, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(s string, direction, x, y int) bool {
	for ; true; {
		if len(s) == 0 {
			if x == tx && y == ty {
				return true
			} else {
				return false
			}
		}
		if s[0] == F {
			switch direction {
			case Right:
				x += 1
			case Up:
				y += 1
			case Left:
				x -= 1
			case Down:
				y -= 1
			}
			s = s[1:]
		} else {
			if slength > len(s) && prefixes[slength-len(s)-1] < (abs(x-tx)+abs(y-ty)) {
				return false
			}
			if len(s) > 1 && s[1] == T {
				a := direction + 2
				if a > 4 {
					a = a - 4
				}
				return dfs(s[2:], direction, x, y) || dfs(s[2:], a, x, y)
			}
			a, b := direction+1, direction+3
			if a > 4 {
				a = a - 4
			}
			if b > 4 {
				b = b - 4
			}
			return dfs(s[1:], a, x, y) || dfs(s[1:], b, x, y)
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs1(s string, direction, x, y, tx, ty int) bool {
	for ; true; {
		if len(s) == 0 {
			if x == tx && y == ty {
				return true
			} else {
				return false
			}
		}
		if s[0] == F {
			switch direction {
			case Right:
				x += 1
			case Up:
				y += 1
			case Left:
				x -= 1
			case Down:
				y -= 1
			}
			s = s[1:]
		} else {
			a, b := direction+1, direction+3
			if a > 4 {
				a = a - 4
			}
			if b > 4 {
				b = b - 4
			}
			return dfs1(s[1:], a, x, y, tx, ty) || dfs1(s[1:], b, x, y, tx, ty)
		}
	}
	return false
}
