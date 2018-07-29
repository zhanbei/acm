package main

import (
	"fmt"
	"strings"
)

// 2018-07-29T20:36:21+0800
func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(strings.Replace(s, "2017", "2018", -1))
}
