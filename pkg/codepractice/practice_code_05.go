package codepractice

/*
206.反转链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode

	cur := head
	for cur != nil {

		// 记录下个节点
		next := cur.Next

		// 指向上一个节点
		cur.Next = prev

		// prev往后移动一位
		prev = cur

		// cur往后移动一位
		cur = next
	}

	return prev
}

func HandlePratice05Entry() {
	reverseList(&ListNode{})
}
