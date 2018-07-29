package main

import "fmt"

// 2018-07-29T20:49:12+0800
// 2018-07-29T20:55:06+0800
func main() {
	var N, t int
	fmt.Scanf("%d", &N)
	m := make(map[int]bool)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &t)
		m[t] = true
	}
	fmt.Println(len(m))
}
