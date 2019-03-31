package main

import "fmt"

// 2019-03-31T23:08:18+0800
// 2019-03-31T23:14:29+0800
// 2019-03-31T23:22:56+0800 # AC @see https://atcoder.jp/contests/arc079/submissions/4790234
func main() {
	var N, M int
	fmt.Scanf("%d%d", &N, &M)
	//a := make([]int, M)
	//b := make([]int, M)
	var a, b int
	starts := make(map[int]bool, M)
	ends := make(map[int]bool, M)
	for i := 0; i < M; i++ {
		//fmt.Scanf("%d%d", &a[i], &b[i])
		fmt.Scanf("%d%d", &a, &b)
		if a == 1 {
			starts[b] = true
		} else if b == 1 {
			starts[a] = true
		} else if a == N {
			ends[b] = true
		} else if b == N {
			ends[a] = true
		}
	}

	for v := range starts {
		if ends[v] {
			fmt.Println("POSSIBLE")
			return
		}
	}
	fmt.Println("IMPOSSIBLE")
}
