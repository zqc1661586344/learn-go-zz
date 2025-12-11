package codepractice

import (
	"fmt"
	"strings"
)

/*
557. 反转字符串中的单词 III
*/

func reverseWords(s string) string {

	sSlices := strings.Split(s, " ")
	fmt.Printf("the sSlice is: %v\n", sSlices)

	length := len(sSlices)
	res := ""

	for i := 0; i < length; i++ {
		res = res + handleReverseSingleWord(sSlices[i]) + " "
	}

	res = strings.Trim(res, " ")
	fmt.Printf("the res is: %v\n", res)
	return res
}

func handleReverseSingleWord(s string) string {
	byteSlice := []byte(s)
	length := len(byteSlice)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		byteSlice[i], byteSlice[j] = byteSlice[j], byteSlice[i]
	}

	return string(byteSlice)
}

func HandlePratice04Entry() {
	//     输入：s = "Let's take LeetCode contest"
	// 输出："s'teL ekat edoCteeL tsetnoc"
	// 输入： s = "Mr Ding"
	// 输出："rM gniD"
	s := "Let's take LeetCode contest"
	s = "Mr Ding"
	fmt.Println(reverseWords(s))
}
