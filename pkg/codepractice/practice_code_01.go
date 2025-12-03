package codepractice

import (
	"fmt"
)

func compareMap(map1, map2 map[string]int) bool {

	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		value2, ok := map2[key]
		if ok {
			if value1 != value2 {
				return false
			}
			continue
		}
		return false
	}

	return true
}

func intInSlice(num int, nums []int) bool {
	for _, j := range nums {
		if j == num {
			return true
		}
	}
	return false
}

func countStrNums(str string) map[string]int {
	conutMap := make(map[string]int)
	for _, v := range []rune(str) {
		fmt.Println(string(v))
		_, ok := conutMap[string(v)]
		if ok {
			conutMap[string(v)] += 1
		} else {
			conutMap[string(v)] = 1
		}
	}
	fmt.Printf("the str is: %v, thr countMap is: %+v\n", str, conutMap)
	return conutMap
}

func groupAnagrams(strs []string) [][]string {
	// 统计每个字符串中字母的数目
	staMapS := make([]map[string]int, 0, len(strs))
	for _, value := range strs {
		staMapS = append(staMapS, countStrNums(value))
	}
	fmt.Printf("the staMapS is: %v\n", staMapS)

	// 从第一个元素开始搜索异位词，并记录index
	indexStas := make([][]int, 0)
	alreadyCount := make([]int, 0, len(strs))
	for i := 0; i < len(strs); i++ {
		sta := make([]int, 0)

		if intInSlice(i, alreadyCount) {
			continue
		}

		sta = append(sta, i)
		alreadyCount = append(alreadyCount, i)
		fmt.Printf("the sta 1 is: %v\n", sta)

		for j := i + 1; j < len(strs); j++ {
			if compareMap(staMapS[i], staMapS[j]) {
				sta = append(sta, j)
				alreadyCount = append(alreadyCount, j)
			}
		}

		indexStas = append(indexStas, sta)

	}

	fmt.Printf("the indexStas is: %v\n", indexStas)

	result := make([][]string, 0, len(indexStas))

	for _, indexs := range indexStas {
		tmp := make([]string, 0, len(indexs))
		for _, index := range indexs {
			tmp = append(tmp, strs[index])
		}
		result = append(result, tmp)
	}

	fmt.Printf("the result is: %+v\n", result)
	return result
}

func HandlePratice01Entry() {
	strs2 := []string{"ddddddddddg", "dgggggggggg"}
	groupAnagrams(strs2)
}
