package main

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

type ArrayQueue struct {
	Q     []int
	front int
	rear  int
	count int
}

func (q *ArrayQueue) Enqueue(a int) {
	if q.IsEmpty() {
		q.Q = append(q.Q, a)
		q.count++
		return
	} else {
		q.Q = append(q.Q, a)
		q.rear++
		q.count++
	}
}

func (q *ArrayQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("An empty queue")
	}
	removed := q.Q[q.front]
	if q.front < q.rear {
		q.Q = q.Q[q.front+1:]
		q.rear--
		q.count--
	} else if q.front == q.rear {
		q.Q = nil
		q.count--
	}
	return removed
}

func (q *ArrayQueue) IsEmpty() bool {
	return len(q.Q) == 0
}

func main() {
	var q ArrayQueue
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.ReverseK(3)

}

func (q ArrayQueue) ReverseK(k int) {
	var s Stack
	// Deque K elements to a stack
	for i := 0; i < k; i++ {
		s.Push(q.Dequeue())
	}

	// Add the stack elements at the end of the queue
	for i := 0; i < k; i++ {
		a, _ := s.Pop()
		q.Enqueue(a)
	}

	for i := 0; i < len(q.Q)-k; i++ {
		q.Enqueue(q.Dequeue())
	}
	fmt.Println(q)

}

// reversek(3)
// [10, 20, 30, 40, 50] => [30, 20, 10, 40, 50]
