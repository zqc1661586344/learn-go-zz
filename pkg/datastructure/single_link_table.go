package datastructure

import (
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type SingleLinkNode[T any] struct {
	head   *Node[T]
	length int
}

// NewSingleLinkNode 构造函数
func NewSingleLinkNode[T any]() *SingleLinkNode[T] {
	return &SingleLinkNode[T]{}
}

// IsEmpty 是否为空
func (s *SingleLinkNode[T]) IsEmpty() bool {
	return s.head == nil
}

// InsertValueToTail 在尾部插入数据
func (s *SingleLinkNode[T]) InsertValueToTail(value T) {
	newNode := &Node[T]{Value: value}

	if s.IsEmpty() {
		s.head = newNode
		s.length++
		return
	}

	cur := s.head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = newNode
	s.length++
}

// InsertValueToHead 在头部插入数据
func (s *SingleLinkNode[T]) InsertValueToHead(value T) {
	newNode := &Node[T]{Value: value}

	cur := s.head
	s.head = newNode
	s.length++
	newNode.Next = cur
}

// Traverse 遍历
func (s *SingleLinkNode[T]) Traverse() []T {
	res := make([]T, 0, s.length)

	if s.IsEmpty() {
		fmt.Println("empty single link node")
		return res
	}

	cur := s.head
	for cur != nil {
		fmt.Printf("the node is: %+v\n", cur.Value)
		res = append(res, cur.Value)
		cur = cur.Next
	}

	fmt.Printf("the all link node value is: %+v\n", res)
	return res
}

// RemoveItemUseIndex 删除某位置的值
func (s *SingleLinkNode[T]) RemoveItemUseIndex(index int) (value T, err error) {
	if index < 0 || index >= s.length {
		err = fmt.Errorf("illeagel index")
		return
	}

	if s.IsEmpty() {
		err = fmt.Errorf("empty link node")
		return
	}

	cur := s.head
	for range index {
		if cur != nil {
			cur = cur.Next
		}
	}

	delNode := cur.Next
	cur.Next = delNode.Next
	s.length--
	return delNode.Value, nil
}

// 在某个位置添加值
func (s *SingleLinkNode[T]) InsertItemUseIndex(index int, value T) {
	if index < 0 || index >= s.length {
		fmt.Printf("illeagel index: %v for insert\n", index)
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	if s.IsEmpty() && index == 0 {
		s.head = newNode
		s.length++
		return
	}

	cur := s.head
	for range index {
		if cur.Next != nil {
			cur = cur.Next
		}
	}

	tmp := cur.Next
	cur.Next = newNode
	newNode.Next = tmp

	s.length++
}

// 反转 Reverse
func (s *SingleLinkNode[T]) Reverse() {
	var pre *Node[T]
	cur := s.head

	for cur != nil {
		// 记录当前的下一个节点
		next := cur.Next

		// 当前的next为上一个节点
		cur.Next = pre

		// 上一个节点切换到当前节点
		pre = cur

		// 当前节点切换到下一个
		cur = next
	}

	s.head = pre
}

func HandleSingleLinkTableEntry() {
	s := NewSingleLinkNode[int]()

	s.InsertValueToTail(2)
	s.InsertValueToTail(3)
	res1 := s.Traverse()
	fmt.Printf("the link node is: %v, the lenth is: %v, the values is: %v\n", s, s.length, res1)
}
