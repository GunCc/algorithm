package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {

	left, right := 0, len(nums)
	for left < right {
		fmt.Println("left=>", left, "  right=>", right)
		if nums[left] == val {
			fmt.Println("nums[left]", nums[left], "  nums[right-1]", nums[right-1])
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	for _, v := range nums {
		fmt.Println("v", v)
	}
	return left

	// left := 0
	// for _, v := range nums { // v 即 nums[right]
	//     if v != val {
	//         nums[left] = v
	//         left++
	//     }
	// }
	// return left

}
