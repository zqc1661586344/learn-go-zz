package datastructure

import "fmt"

/*
二叉搜索树
DFS（基于深度遍历）
前序遍历：根左右
中序遍历：左根右
后序遍历：左右根

BFS（基于广度遍历）

特点：
1、中序遍历得到的是一个升序的排列
2、左<根<右
*/

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

type BSTree struct {
	Root *TreeNode
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

func NewBSTreeWithRoot(value int) *BSTree {
	return &BSTree{Root: NewTreeNode(value)}
}

// -----------递归写法-----------
func insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return NewTreeNode(value)
	}

	// 插入左树
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value { // 插入右树
		node.Right = insertNode(node.Right, value)
	} else {
		fmt.Printf("the value: %v is exist\n", value)
	}

	return node
}

// 插入节点
func (s *BSTree) InsertNode(value int) {
	s.Root = insertNode(s.Root, value)
}

// 前序遍历：根左右
func preTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Value)
	*values = append(*values, node.Value)
	preTraversal(node.Left, values)
	preTraversal(node.Right, values)
}

func (s *BSTree) PreTraversal() []int {
	result := []int{}
	if s == nil || s.Root == nil {
		return result
	}

	preTraversal(s.Root, &result)
	return result
}

// 中序遍历：左根右
func inTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	inTraversal(node.Left, values)
	fmt.Printf("%d ", node.Value)
	*values = append(*values, node.Value)
	inTraversal(node.Right, values)
}

func (s *BSTree) InTraversal() []int {
	result := []int{}
	if s == nil || s.Root == nil {
		return result
	}

	inTraversal(s.Root, &result)
	return result
}

// 后序遍历：左右根
func postTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	postTraversal(node.Left, values)
	postTraversal(node.Right, values)
	fmt.Printf("%d ", node.Value)
	*values = append(*values, node.Value)
}

func (s *BSTree) PostTraversal() []int {
	result := []int{}
	if s == nil || s.Root == nil {
		return result
	}

	postTraversal(s.Root, &result)
	return result
}

// -----------递归写法-----------

// -----------迭代写法-----------
func (s *BSTree) InsertNodeUseLoop(value int) {
	if s == nil {
		return
	}

	newNode := NewTreeNode(value)

	if s.Root == nil {
		s.Root = newNode
		return
	}

	cur := s.Root

	// // 写法一
	// for cur != nil {
	//  if value < cur.Value {
	//   if cur.Left == nil {
	//    cur.Left = newNode
	//    return
	//   }

	//   cur = cur.Left
	//  } else if value > cur.Value {
	//   if cur.Right == nil {
	//    cur.Right = newNode
	//    return
	//   }

	//   cur = cur.Right
	//  } else {
	//   fmt.Printf("the value: %v is exist\n", value)
	//   return
	//  }
	// }

	// 写法二
	// 记录最后一个节点
	var par *TreeNode
	for cur != nil {
		par = cur

		if value < cur.Value {
			cur = cur.Left
		} else if value > cur.Value {
			cur = cur.Right
		} else {
			fmt.Printf("the value: %v is exist\n", value)
			return
		}
	}

	if value < par.Value {
		par.Left = newNode
	} else if value > par.Value {
		par.Right = newNode
	} else {
		fmt.Printf("the value: %v is exist\n", value)
	}
}

// 前序遍历：根左右
func (s *BSTree) PreTraversalUseLoop() []int {
	result := []int{}
	if s == nil || s.Root == nil {
		return result
	}

	stack := []*TreeNode{s.Root}
	for len(stack) > 0 {

		// 取出栈顶元素
		node := stack[len(stack)-1]
		fmt.Printf("%d ", node.Value)
		result = append(result, node.Value)

		// 弹出栈顶元素
		stack = stack[:len(stack)-1]

		// 先压右树元素
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// 再压左树元素
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// 中序遍历：左根右
func (s *BSTree) InTraversalUseLoop() []int {
	result := []int{}
	if s == nil || s.Root == nil {
		return result
	}

	cur := s.Root
	stack := []*TreeNode{}

	for cur != nil || len(stack) > 0 {
		// 先处理左子树
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 取出栈顶元素
		cur = stack[len(stack)-1]
		fmt.Printf("%d ", cur.Value)
		result = append(result, cur.Value)

		// 弹出栈顶元素
		stack = stack[:len(stack)-1]

		cur = cur.Right
	}

	return result
}

// 后序遍历：左右根
func (s *BSTree) PostTraversalUseLoop() []int {
	result := []int{}
	if s == nil || s.Root == nil {
		return result
	}

	// 栈保存待处理的节点，初始化时传入Root指向的节点
	stack1 := []*TreeNode{s.Root}
	stack2 := []*TreeNode{}

	for len(stack1) > 0 {
		// 获取栈顶元素
		node := stack1[len(stack1)-1]

		// 弹栈
		stack1 = stack1[:len(stack1)-1]

		// 将栈顶元素压入第二个栈
		stack2 = append(stack2, node)

		// 压入左子树（保证弹出顺序按照左--->右）
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}

		// 压入右子树
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	// 遍历第2个栈得到结果
	for len(stack2) > 0 {
		top := stack2[len(stack2)-1]
		fmt.Printf("%d ", top.Value)
		result = append(result, top.Value)

		stack2 = stack2[:len(stack2)-1]
	}

	return result
}

// -----------迭代写法-----------

func HandleBinarySearchTreeEntry() {
	tree := NewBSTree()

	// 递归插入元素
	tree.InsertNode(4)
	tree.InsertNode(2)
	tree.InsertNode(5)
	tree.InsertNode(6)
	tree.InsertNode(10)
	tree.InsertNode(9)

	// 遍历插入元素
	// tree.InsertNodeUseLoop(4)
	// tree.InsertNodeUseLoop(2)
	// tree.InsertNodeUseLoop(5)
	// tree.InsertNodeUseLoop(6)

	// fmt.Printf("前序遍历: %v\n", tree.PreTraversal())
	// fmt.Printf("前序遍历： %v\n", tree.PreTraversalUseLoop())
	fmt.Printf("中序遍历: %v\n", tree.InTraversal())
	fmt.Printf("中序遍历: %v\n", tree.InTraversalUseLoop())
	fmt.Printf("后序遍历: %v\n", tree.PostTraversal())
	fmt.Printf("后序遍历: %v\n", tree.PostTraversalUseLoop())
}
