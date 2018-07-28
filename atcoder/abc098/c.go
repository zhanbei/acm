package main

import (
	"fmt"
	"math"
)

const MaxN = 300005

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	var i int
	var s string
	fmt.Scanf("%d", &i)
	fmt.Scanf("%s", &s)
	// prefix sum, @see https://en.wikipedia.org/wiki/Prefix_sum
	// i -> how many 'W' of s[0:i+1].
	var prefixes [MaxN]int
	if s[0] == 'W' {
		prefixes[0] = 1
	} else {
		prefixes[0] = 0
	}
	for i := 1; i < len(s); i++ {
		if s[i] == 'W' {
			prefixes[i] = prefixes[i-1] + 1
		} else {
			prefixes[i] = prefixes[i-1]
		}
	}
	n := math.MaxInt64
	for i := 0; i < len(s); i++ {
		// The altered s must be 'EEEEEEOWWWWWWW'.
		// Sum of 'W' in left and 'E' in the right;
		if s[i] == 'W' {
			n = min(n, prefixes[i]-1+(len(s)-(i+1))-(prefixes[len(s)-1]-prefixes[i]))
		} else {
			n = min(n, prefixes[i]+(len(s)-(i+1))-(prefixes[len(s)-1]-prefixes[i]))
		}
	}
	fmt.Println(n)
}
