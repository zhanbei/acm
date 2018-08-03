package main

import (
	"fmt"
	"strings"
)

// 2018-08-03T23:01:50+0800
// 2018-08-03T23:12:52+0800 # @see https://arc063.contest.atcoder.jp/submissions/2939366
func main() {
	var S string
	fmt.Scanf("%s", &S)
	n := len(S)
	for ; true; {
		S = strings.Replace(S, "WW", "W", -1)
		S = strings.Replace(S, "BB", "B", -1)
		if n == len(S) {
			break
		}
		n = len(S)
	}
	switch len(S) {
	case 1:
		fmt.Println(0)
	case 2:
		fmt.Println(1)
	default:
		fmt.Println(len(S) - 1)
	}
}
