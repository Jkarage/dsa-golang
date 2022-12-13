package queueII

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() int {
	i := len(*s) - 1
	popped := (*s)[i]
	*s = (*s)[:i]
	return popped
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Peek() int {
	i := len(*s) - 1
	return (*s)[i]
}

type Queue struct {
	stack1 Stack
	stack2 Stack
}

func (q *Queue) Enqueue(i int) {
	q.stack1.Push(i)
}

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		panic("An empty Queue")
	}
	q.moveToStack2()
	q.stack2.Pop()
}

func (q Queue) IsEmpty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}

func (q Queue) Peek() int {
	if q.IsEmpty() {
		panic("an empty stack")
	}
	q.moveToStack2()

	return q.stack2.Peek()
}

func (q Queue) moveToStack2() {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			popped := q.stack1.Pop()
			q.stack2.Push(popped)
		}
	}
}
