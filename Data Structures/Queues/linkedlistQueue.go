package main

import "fmt"

type Linkedlist struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Value int
	Next  *Node
}

func (l *Linkedlist) AddHead(value int) {
	node := &Node{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.Size++
}

func (l *Linkedlist) AddTail(value int) {
	node := &Node{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size++
}

func (l *Linkedlist) DeleteHead() {
	if l.IsEmpty() {
		panic("Emptying an empty LinkedList")
	} else if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		nextHeader := l.Head.Next
		l.Head.Next = nil
		l.Head = nextHeader
	}
	l.Size--
}

func (l *Linkedlist) DeleteTail() {
	if l.IsEmpty() {
		panic("Dealing with an empty LinkedList")
	}

	if l.Head == l.Tail {
		panic("Dealing with a single node LinkedList")
	}

	previous := l.Previous(l.Tail)
	previous.Next = nil
	l.Tail = previous
	l.Size--
}

func (l *Linkedlist) IndexOf(value int) int {
	index := 0
	current := l.Head
	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

func (l *Linkedlist) Contains(value int) bool {
	return l.IndexOf(value) != -1
}

func (l *Linkedlist) Previous(node *Node) *Node {
	current := l.Head
	for current != nil {
		if current.Next == node {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *Linkedlist) IsEmpty() bool {
	return l.Head == nil
}

func (l *Linkedlist) ToSlice() []int {
	slice := make([]int, l.Size)
	index := 0
	current := l.Head
	for current != nil {
		slice[index] = current.Value
		index++
		current = current.Next
	}
	return slice
}

func (l *Linkedlist) Reverse() {
	p := l.Head
	c := l.Head.Next
	l.Tail = p
	l.Tail.Next = nil
	for c != nil {
		n := c.Next
		c.Next = p
		p = c
		c = n

	}
	l.Head = p
}
func (l *Linkedlist) KthFromEnd(k int) int {
	if l.IsEmpty() {
		panic("An empty linked list")
	}
	left := l.Head
	right := l.Head

	for i := 0; i < k-1; i++ {
		right = right.Next
		if right == nil {
			panic("Illegal argument")
		}
	}

	for right != l.Tail {
		left = left.Next
		right = right.Next
	}
	return left.Value
}

func (l *Linkedlist) PrintMiddle() {
	left := l.Head
	right := l.Head
	for right != l.Tail && right.Next != l.Tail {
		right = right.Next.Next
		left = left.Next
	}

	if right == l.Tail {
		fmt.Println(left.Value)
	} else {
		fmt.Println(left.Value, left.Next.Value)
	}
}

func (l *Linkedlist) HasLoop() bool {
	left := l.Head
	right := l.Head

	for right != nil && right.Next != nil {
		right = right.Next.Next
		left = left.Next

		if left == right {
			return true
		}
	}
	return false
}

type Queue struct {
	L Linkedlist
}

func (q *Queue) Enqueue(a int) {
	q.L.AddTail(a)
}

func (q *Queue) Dequeue() {
	q.L.DeleteHead()
}

func (q *Queue) Peek() int {
	return q.L.Head.Value
}

func (q *Queue) IsEmpty() bool {
	return q.L.IsEmpty()
}
