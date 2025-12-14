package codepractice

import (
	"fmt"
)

/**
2. 两数相加
*/

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func Reverse(table *ListNode) *ListNode {
	if table == nil {
		return nil
	}

	var prev *ListNode
	cur := table
	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}

func Travel(table *ListNode) string {
	s := ""
	cur := table
	for cur != nil {
		s += fmt.Sprintf("%d", cur.Val)
		cur = cur.Next
	}

	fmt.Printf("the travel result is: %v\n", s)
	return s
}

func ApendTail(nums []int) *ListNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}

	cur := head
	for i := 1; i < length; i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	fmt.Printf("the res is: %v\n", head)
	return head
}

func addTwoNumbers(l1Rev *ListNode, l2Rev *ListNode) *ListNode {
	// 统计各位数相加的结果
	sumSli := []int{}
	carry := 0
	for l1Rev != nil && l2Rev != nil {
		tmp := 0
		if l1Rev.Val+l2Rev.Val+carry >= 10 {
			tmp = (l1Rev.Val + l2Rev.Val + carry) - 10
			carry = 1
		} else {
			tmp = l1Rev.Val + l2Rev.Val + carry
			carry = 0
		}
		sumSli = append(sumSli, tmp)

		l1Rev = l1Rev.Next
		l2Rev = l2Rev.Next
	}

	fmt.Printf("the first sum: %v, the carry is: %v\n", sumSli, carry)
	fmt.Printf("the l1rev is: %v, the l2rev is: %v\n", Travel(l1Rev), Travel(l2Rev))

	// 结束
	if l1Rev == nil && l2Rev == nil {
		if carry != 0 {
			sumSli = append(sumSli, carry)
		}

		fmt.Printf("the l1 and l2 is equal, the final sum is: %v\n", sumSli)
		// 构造链表
		resTable := ApendTail(sumSli)
		return resTable
	}

	var cur *ListNode
	if l1Rev == nil {
		fmt.Println("---l2---")
		// l1比l2短
		cur = l2Rev
	} else {
		fmt.Println("---l1---")
		// l2比l1短
		cur = l1Rev
	}

	for cur != nil {
		tmp := 0
		if cur.Val+carry >= 10 {
			tmp = (cur.Val + carry) - 10
			carry = 1
		} else {
			tmp = cur.Val + carry
			carry = 0
		}
		sumSli = append(sumSli, tmp)
		cur = cur.Next
	}

	fmt.Printf("the second sum: %v, the carry is: %v\n", sumSli, carry)

	if carry != 0 {
		sumSli = append(sumSli, carry)
	}

	fmt.Printf("the final sum is: %v\n", sumSli)

	// 构造链表
	resTable := ApendTail(sumSli)
	// // 反转链表
	// resTable = Reverse(resTable)
	Travel(resTable)

	return resTable
}

// 下面的这种写法，把string转化成int的方式，在数据很大时越界了，可能需要使用大数计算
// //////////////////////

// func Travel(table *ListNode) string {
// 	s := ""
// 	cur := table
// 	for cur != nil {
// 		s += fmt.Sprintf("%d", cur.Val)
// 		cur = cur.Next
// 	}

// 	fmt.Printf("the travel result is: %v\n", s)
// 	return s
// }

// func ApendTail(nums []int) *ListNode {
// 	length := len(nums)
// 	if length == 0 {
// 		return nil
// 	}

// 	head := &ListNode{Val: nums[0]}

// 	cur := head
// 	for i := 1; i < length; i++ {
// 		cur.Next = &ListNode{Val: nums[i]}
// 		cur = cur.Next
// 	}

// 	fmt.Printf("the res is: %v\n", head)
// 	return head
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// l1Str := Travel(l1)
// 	// l2Str := Travel(l2)

// 	// l1,le先反转
// 	l1Rev := Reverse(l1)
// 	l2Rev := Reverse(l2)

// 	// 获取l1和l2数据
// 	l1Str := Travel(l1Rev)
// 	l2Str := Travel(l2Rev)

// 	l1int, err := strconv.Atoi(l1Str)
// 	if err != nil {
// 		return nil
// 	}

// 	l2int, err := strconv.Atoi(l2Str)
// 	if err != nil {
// 		return nil
// 	}

// 	fmt.Printf("the l1 is: %v, the l2 is: %v\n", l1int, l2int)
// 	sum := l1int + l2int
// 	fmt.Printf("the sum is: %v\n", sum)

// 	strSum := strconv.Itoa(sum)

// 	intNum := []int{}
// 	for _, num := range strSum {
// 		n, _ := strconv.Atoi(string(num))
// 		intNum = append(intNum, n)
// 	}
// 	fmt.Printf("the intNum is: %v\n", intNum)

// 	// 构造链表
// 	resTable := ApendTail(intNum)

// 	// 反转链表
// 	resTable = Reverse(resTable)
// 	Travel(resTable)

// 	return resTable
// }
// //////////////////////

func HandlePratice12Entry() {
	// 输入：l1 = [2,4,3], l2 = [5,6,4]
	// 输出：[7,0,8]
	// testTable1 := &ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val: 4,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }

	// testTable2 := &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val: 6,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }

	// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	// 输出：[8,9,9,9,0,0,0,1]
	// testTable1 := &ListNode{
	// 	Val: 9,
	// 	Next: &ListNode{
	// 		Val: 9,
	// 		Next: &ListNode{
	// 			Val: 9,
	// 			Next: &ListNode{
	// 				Val: 9,
	// 				Next: &ListNode{
	// 					Val: 9,
	// 					Next: &ListNode{
	// 						Val: 9,
	// 						Next: &ListNode{
	// 							Val: 9,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// testTable2 := &ListNode{
	// 	Val: 9,
	// 	Next: &ListNode{
	// 		Val: 9,
	// 		Next: &ListNode{
	// 			Val: 9,
	// 			Next: &ListNode{
	// 				Val: 9,
	// 			},
	// 		},
	// 	},
	// }

	// 输入：l1 = [0], l2 = [0]
	// 输出：[0]
	// testTable1 := &ListNode{Val: 0}
	// testTable2 := &ListNode{Val: 0}

	// 输入：l1 = [2,4,9], l2 = [5,6,4,9]
	// 输出：[7,0,4,0,1]
	// testTable1 := &ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val: 4,
	// 		Next: &ListNode{
	// 			Val: 9,
	// 		},
	// 	},
	// }
	// testTable2 := &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val: 6,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 			Next: &ListNode{
	// 				Val: 9,
	// 			},
	// 		},
	// 	},
	// }

	// 输入：l1 = [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1], l2 = [5,6,4]
	// 输出：[6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]

	// 输入：l1 = [5], l2 = [5]
	// 输出：[0,1]
	testTable1 := &ListNode{Val: 5}
	testTable2 := &ListNode{Val: 5}

	// Travel(testTable1)

	// rev := Reverse(testTable1)
	// Travel(rev)
	// // fmt.Printf("the rev 1 is: %v\n", Travel(rev))
	// fmt.Println("-----------------------")

	// newTable := ApendTail([]int{7, 8, 9})
	// Travel(newTable)
	fmt.Println("********************")

	addTwoNumbers(testTable1, testTable2)
}
