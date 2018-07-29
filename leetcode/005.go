package main

import "fmt"

// 2018-06-30T15:09:32+0800
func main() {
	out := longestPalindrome("helloworldmxxmdl")
	out = longestPalindrome("cbbd")
	fmt.Println(out)
}

// 2018-06-30T15:43:53+0800
func longestPalindrome(s string) string {
	l := 0
	ans := ""
	for start := 0; start < len(s); start++ {
		for i := 1; ; i++ {
			if start-i < 0 || start+i >= len(s) || s[start-i] != s[start+i] {
				if 2*i-1 > l {
					l = 2*i - 1
					// start-i+1 : start+i-1
					ans = s[start-i+1 : start+i]
					//if l != len(ans) {
					//	fmt.Println(l, len(ans), ans)
					//}
				}
				break
			}
		}
		for i := 0; ; i++ {
			if start-i < 0 || start+i+1 >= len(s) || s[start-i] != s[start+i+1] {
				if 2*i > l {
					l = 2 * i
					// start-i+1 : start+i
					ans = s[start-i+1 : start+i+1]
					//if l != len(ans) {
					//	fmt.Println(l, len(ans), ans)
					//}
				}
				break
			}
		}
	}
	return ans
}
