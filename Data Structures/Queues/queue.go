package queue

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

func (q *ArrayQueue) Dequeue() {
	if q.IsEmpty() {
		panic("An empty queue")
	}
	if q.front < q.rear {
		q.Q = q.Q[q.front+1:]
		q.rear--
		q.count--
	} else if q.front == q.rear {
		q.Q = nil
		q.count--
	}
}

func (q *ArrayQueue) IsEmpty() bool {
	return len(q.Q) == 0
}
