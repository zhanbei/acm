package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var s string
	var i int
	fmt.Scanf("%d", &i)
	fmt.Scanf("%s", &s)
	n := 0
	for i := 1; i < len(s); i++ {
		n = max(n, solve(s[0:i], s[i:]))
	}
	fmt.Println(n)
}

func solve(a, b string) int {
	ma := make(map[uint8]bool)
	mb := make(map[uint8]bool)
	for i := 0; i < len(a); i++ {
		ma[a[i]] = true
	}
	for i := 0; i < len(b); i++ {
		mb[b[i]] = true
	}
	res := 0
	for key := range ma {
		if mb[key] {
			res++
		}
	}
	return res
}
