package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println("nums1", nums1)
}
func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
