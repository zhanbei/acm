package main

import "fmt"

func main() {
	out := findMedianSortedArrays([]int{}, []int{2})
	fmt.Println(out)
}

// 2018-06-30T15:09:13+0800
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := nums1
	n := nums2
	ml := 0
	mr := len(m) - 1
	mb := false
	nl := 0
	nr := len(n) - 1
	nb := false
	if ml > mr {
		mb = true
	}
	if nl > nr {
		nb = true
	}
	// i is the number of numbers passed.
	for i := 0; ; i += 2 {
		if len(m)+len(n)-i <= 2 {
			break
		}
		if nb {
			ml++
		} else if mb {
			nl ++
		} else if m[ml] < n[nl] {
			ml++
		} else {
			nl++
		}
		if ml > mr {
			mb = true
		}
		if nl > nr {
			nb = true
		}
		if nb {
			mr --
		} else if mb {
			nr --
		} else if m[mr] > n[nr] {
			mr--
		} else {
			nr--
		}
		if ml > mr {
			mb = true
		}
		if nl > nr {
			nb = true
		}
	}
	if mb {
		return float64(n[nl]+n[nr]) / 2
	} else if nb {
		return float64(m[ml]+m[mr]) / 2
	} else {
		return float64(m[ml]+n[nl]) / 2
	}
}
