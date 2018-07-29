package main

import (
	"fmt"
)

const MaxRC = 105
const MaxW = 30

// 2018-07-29T14:52:25+0800
// 2018-07-29T18:47:29+0800
func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ii := 0; ii < T; ii++ {
		solve(ii)
	}
}

func solve(caseIndex int) {
	var R, C, W int
	fmt.Scanf("%d%d%d", &R, &C, &W)
	var _strs [MaxRC]string
	var _grid [MaxRC][MaxRC]uint8
	for i := 0; i < R; i++ {
		fmt.Scanf("%s", &_strs[i])
		for ii := 0; ii < C; ii++ {
			_grid[i][ii] = _strs[i][ii]
		}
	}
	var _words [MaxW]string
	//var words [MaxW]uint8
	words := make(map[uint8]bool)
	for i := 0; i < W; i++ {
		fmt.Scanf("%s", &_words[i])
		words[_words[i][0]] = true
	}

	var prefixes [MaxRC][MaxRC]int
	var grid [MaxRC][MaxRC]bool
	// Build the prefixes
	for i := 0; i < R; i++ {
		for ii := 0; ii < C; ii++ {
			grid[i][ii] = words[_grid[i][ii]]
			if i <= 0 && ii <= 0 {
				if grid[i][ii] {
					prefixes[i][ii] = 1
				} else {
					prefixes[i][ii] = 0
				}
			} else if i <= 0 {
				if grid[i][ii] {
					prefixes[i][ii] = prefixes[i][ii-1] + 1
				} else {
					prefixes[i][ii] = prefixes[i][ii-1]
				}
			} else if ii <= 0 {
				if grid[i][ii] {
					prefixes[i][ii] = prefixes[i-1][ii] + 1
				} else {
					prefixes[i][ii] = prefixes[i-1][ii]
				}
			} else {
				if grid[i][ii] {
					prefixes[i][ii] = prefixes[i-1][ii] + prefixes[i][ii-1] - prefixes[i-1][ii-1] + 1
				} else {
					prefixes[i][ii] = prefixes[i-1][ii] + prefixes[i][ii-1] - prefixes[i-1][ii-1]
				}
			}
		}
	}

	var n float64
	nn := make(map[float64][]*Rect)
	//noccur := make(map[float64]int)
	//nlength := make(map[float64]int)
	// Loop for all sub grids.
	for i := 0; i < R; i++ {
		for ii := 0; ii < C; ii++ {
			for k := i; k < R; k++ {
				for kk := ii; kk < C; kk++ {
					// Sub: i, ii -> k, kk
					rect := getScore(&prefixes, i, ii, k, kk)
					nn[rect.score] = append(nn[rect.score], rect)
					if rect.length <= 0 {
						fmt.Println("---->>>", rect.score, rect.occur, rect.length)
					}
					if rect.score >= n {
						n = rect.score
					}
					//noccur[score] = occur
					//nlength[score] = length
				}
			}
		}
	}

	occur, length := simplify(nn[n][0].occur*4, nn[n][0].length)
	//for i := 0; i < len(nn[n]); i++ {
	//	if a, b := simplify(nn[n][i].occur*4, nn[n][i].length); occur != a || length != b {
	//		fmt.Println("WARNING: XXXXXXXXXX", occur, length, a, b)
	//	}
	//}

	//fmt.Printf("Case #%d: %d \n", caseIndex+1, len(nn[n]))
	fmt.Printf("Case #%d: %d/%d %d\n", caseIndex+1, occur, length, len(nn[n]))
}

type Rect struct {
	length int
	occur  int
	score  float64
}

func getScore(prefixes *[MaxRC][MaxRC]int, i, ii, k, kk int) *Rect {
	//func getScore(prefixes *[MaxRC][MaxRC]int, i, ii, k, kk int) (int, int, float64) {
	//	return 0, 0, float64(0)
	var a, b, c, d int
	if i <= 0 {
		b = 0
	} else {
		b = prefixes[i-1][kk]
	}
	if ii <= 0 {
		c = 0
	} else {
		c = prefixes[k][ii-1]
	}
	if i <= 0 || ii <= 0 {
		d = 0
	} else {
		d = prefixes[i-1][ii-1]
	}
	a = prefixes[k][kk]
	occur := a - b - c + d
	length := (k - i + 1) + (kk - ii + 1)
	score := float64(occur) / float64(length)
	//return occur, length, score
	if length == 0 {
		length = 1
	}
	return &Rect{length, occur, score}
}

func simplify(a, b int) (int, int) {
	c := getMaxCommon(a, b)
	return a / c, b / c
}

func getMaxCommon(a, b int) int {
	if b == 0 {
		return a
	}
	return getMaxCommon(b, a%b)
}
