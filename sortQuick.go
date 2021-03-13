package algorithms

import (
	"math/rand"
	"time"
)

// SortQuick uses quick sort algorithm to sort an array.
// Different implemenations could be found on the internet.
// I however chose to use the simplest:
// 1. We choose random element x from an array
// 2. Divide array into 3 sub-arrays: left, middle, right
//    left: elements that are less than x
//    middle: elements that are equal to x
//    right: elements that are greater than x
// 3. Then we add 3 recursively sorted arrays.
func SortQuick(a []int) []int {
	if len(a) > 1 {
		rand.Seed(time.Now().Unix())
		var left, mid, right []int
		x := a[rand.Int()%len(a)]
		for i := 0; i < len(a); i++ {
			if a[i] < x {
				left = append(left, a[i])
			} else if a[i] == x {
				mid = append(mid, a[i])
			} else {
				right = append(right, a[i])
			}
		}
		return append(append(SortQuick(left), mid...), SortQuick(right)...)
	}
	return a
}
