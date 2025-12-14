package codepractice

import "fmt"

/*
136. 只出现一次的数字
*/

func singleNumber(nums []int) int {
	fmt.Printf("the nums is: %v\n", nums)

	staMap := make(map[int]int)

	for _, num := range nums {
		if _, ok := staMap[num]; ok {
			staMap[num] += 1
			continue
		}
		staMap[num] = 1
	}
	fmt.Printf("the staMap is: %v\n", staMap)

	for key, value := range staMap {
		if value == 1 {
			return key
		}
	}

	return 0
}

func HandlePratice10Entry() {
	// 输入：nums = [4,1,2,1,2]
	// 输出：4

	// 输入：nums = [2,2,1]
	// 输出：1
	nums := []int{4, 1, 2, 1, 2}
	// nums = []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}
