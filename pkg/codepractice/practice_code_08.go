package codepractice

/*
217.存在重复元素
*/

import (
	"fmt"
)

// containsDuplicate
func containsDuplicate(nums []int) bool {
	staMap := make(map[int]int)

	for _, value := range nums {
		if _, ok := staMap[value]; ok {
			staMap[value] += 1
		} else {
			staMap[value] = 1
		}
	}
	fmt.Printf("the map is: %v\n", staMap)
	sliLen := len(nums)
	mapLen := len(staMap)

	return !(sliLen == mapLen)
}

func HandlePratice08Entry() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums1))

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2))
}
