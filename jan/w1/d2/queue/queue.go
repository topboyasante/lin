package queue

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(val T) {
	q.elements = append(q.elements, val)
}

func (q *Queue[T]) Dequeue() {
	q.elements = q.elements[1:]
}

func (q *Queue[T]) Front() T {
	return q.elements[0]
}

func (q *Queue[T]) IsEmpty() bool {
	if len(q.elements) > 0 {
		return false
	}
	return true
}
