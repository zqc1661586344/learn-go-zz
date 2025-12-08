package codepractice

import (
	"fmt"
)

/*
切片反转
*/

// ReverseSlice 切片反转（直接更新原切片）
func ReverseSlice[T any](slices []T) {
	for i := 0; i < len(slices)/2; i++ {
		j := len(slices) - i - 1
		slices[i], slices[j] = slices[j], slices[i]
	}
}

// ReverseSliceToNewSlice 切片反转（返回新切片）
func ReverseSliceToNewSlice[T any](slices []T) []T {
	res := make([]T, len(slices))
	for i := 0; i < len(slices)/2; i++ {
		j := len(slices) - i - 1

		res[i], res[j] = slices[j], slices[i]
	}

	return res
}

func HandlePratice02Entry() {

	tests := [][]int{
		[]int{},
		[]int{3},
		[]int{4, 2},
		[]int{2, 3, 1},
		[]int{3, 4, 5, 1},
	}

	for i, j := range tests {
		ReverseSlice(j)
		fmt.Println(tests[i])
	}
	fmt.Println("the tests after reverse is: ", tests)

	fmt.Println("###########################")

	for i, j := range tests {
		tmp := ReverseSliceToNewSlice(j)
		fmt.Printf("the tmp is: %+v\n", tmp)
		fmt.Printf("the origin is: %+v\n", tests[i])
	}
	fmt.Printf("the origin tests is: %v\n", tests)
}
