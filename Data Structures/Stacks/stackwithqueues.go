package main

import "fmt"

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

type Stack struct {
	q1  ArrayQueue
	q2  ArrayQueue
	top int
}

func (s *Stack) Push(a int) {
	s.q1.Enqueue(a)
	s.top = a

}

func (s *Stack) Pop() {
	for s.q1.count > 1 {
		s.top = s.q1.Dequeue()
		s.q2.Enqueue(s.top)
	}
	s.swap()
	s.q2.Dequeue()
}

func (s *Stack) swap() {
	s.q1, s.q2 = s.q2, s.q1
}

func (s *Stack) IsEmpty() bool {
	return s.q1.IsEmpty()
}

func (s *Stack) Peek() int {
	return s.top
}

func main() {
	var s Stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)
	s.Pop()
	s.Pop()
	fmt.Println(s.Peek())
}
