package main

import "fmt"

// 2018-07-30T21:14:16+0800
// 2018-07-30T21:16:27+0800
func main() {
	var s string
	fmt.Scanf("%s", &s)
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			n++
		} else if s[i] == '-' {
			n--
		}
	}
	fmt.Println(n)
}
