package main

import (
	"fmt"
)

// 2018-06-30T15:45:28+0800
func main() {
	out := convertToTitle(26)
	//out = convertToTitle(27)
	fmt.Println(out)
}

// 2018-06-30T15:55:48+0800
// 2018-06-30T16:32:17+0800
func convertToTitle(n int) string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var out []uint8
	for i := 0; n != 0; i++ {
		n -= 1
		out = append([]uint8{letters[n%26]}, out...)
		n = n / 26
	}
	return string(out)
}
