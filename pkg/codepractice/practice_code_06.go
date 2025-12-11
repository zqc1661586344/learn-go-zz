package codepractice

import (
	"fmt"
	"sort"
)

/*
88. 合并两个有序数组
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n {
		return
	}

	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	fmt.Printf("the nums1 before: %v\n", nums1)
	sort.Ints(nums1)
	fmt.Printf("the nums1 before: %v\n", nums1)

	// sort.Reverse(sort.IntSlice(nums1))
	// fmt.Printf("the nums1 final: %v\n", nums1)
}

func HandlePratice06Entry() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	// nums1 := []int{1}
	// m := 1
	// nums2 := []int{}
	// n := 0

	// nums1 := make([]int,1)
	// m := 0
	// nums2 := []int{1}
	// n := 1

	merge(nums1, m, nums2, n)
}
