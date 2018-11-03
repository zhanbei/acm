package main

import "fmt"

var S string
var tokens = []string{"dream", "dreamer", "erase", "eraser"}

// 2018-11-03T17:36:05+0800
// 2018-11-03T17:42:17+0800 Try DFS
// 2018-11-03T18:04:01+0800 # @see https://arc065.contest.atcoder.jp/submissions/3524871
func main() {
	fmt.Scanf("%s", &S)
	if dfs(0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func dfs(index int) bool {
	if len(S) == index {
		return true
	}
	if len(S)-index < 5 {
		return false
	}
	if len(S)-index > 7 {
		return (S[index:index+5] == tokens[0] && dfs(index+5)) ||
			(S[index:index+5] == tokens[2] && dfs(index+5)) ||
			(S[index:index+6] == tokens[3] && dfs(index+6)) ||
			(S[index:index+7] == tokens[1] && dfs(index+7))
	}
	if len(S)-index == 5 {
		return S[index:index+5] == tokens[0] || S[index:index+5] == tokens[2]
	}
	if len(S)-index == 6 {
		return S[index:index+6] == tokens[3]
	}
	if len(S)-index == 7 {
		return S[index:index+7] == tokens[1]
	}
	return false
}
