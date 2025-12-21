package codepractice

import "fmt"

/*
230. 二叉搜索树中第 K 小的元素
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左根右
func midTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	// 左
	midTraversal(node.Left, values)

	// 根
	fmt.Printf("%d ", node.Val)
	*values = append(*values, node.Val)

	// 右
	midTraversal(node.Right, values)
}

func kthSmallest(root *TreeNode, k int) int {
	result := []int{}
	midTraversal(root, &result)

	fmt.Printf("the resulrt is: %v\n", result)

	if k > len(result) {
		return 0
	}

	return result[k-1]
}
