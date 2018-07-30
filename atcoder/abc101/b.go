package main

import (
	"fmt"
	"strconv"
)

// 2018-07-30T21:19:57+0800
// 2018-07-30T21:25:19+0800
func main() {
	var N int
	fmt.Scanf("%d", &N)
	sN := strconv.Itoa(N)
	n := 0
	for i := 0; i < len(sN); i++ {
		n += int(sN[i] - 48)
	}
	if N%n == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
