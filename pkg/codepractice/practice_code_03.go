package codepractice

import (
	"fmt"
	"sort"
)

/*
sort包使用：

sort.Ints(a []int)	对int切片升序排序	数值从小到大
sort.Float64s(a []float64)	对float64切片升序排序	数值从小到大（忽略 NaN）
sort.Strings(a []string)	对string切片升序排序	字符串 Unicode 码点字典序
sort.Sort(sort.Reverse(sort.Interface))	对任意实现接口的类型降序排序	基于原 Less 规则反转

sort.Slice：非稳定排序（不保证相等元素的相对位置，底层是快排，效率更高）；
sort.SliceStable：稳定排序（保证相等元素的相对位置，底层是归并排序）。
*/

func HandlePratice03Entry() {
	ints := []int{5, 3, 1, 3, 4, 2}

	sort.Sort(sort.IntSlice(ints))
	fmt.Println("----0", ints)

	// 对int值进行从大到小进行排序
	sort.Ints(ints)
	fmt.Println("----1", ints)

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("----2", ints)

}
