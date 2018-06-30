package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("Hello-world")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]bool)
	l := 0
	// for start := 0; start + l + 1 < len(s); start++ {
	for start := 0; start < len(s); start++ {
		for i := start; ; {
			c := s[i]
			if m[c] {
				if len(m) > l {
					l = len(m)
				}
				m = make(map[uint8]bool)
				break
			} else {
				m[c] = true
			}
			i++
			if i == len(s) {
				if len(m) > l {
					l = len(m)
				}
				m = make(map[uint8]bool)
				break
			}
		}
	}
	return l
}
