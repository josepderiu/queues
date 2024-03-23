package queues

import (
	"fmt"
)

type T interface{}

type Queue struct {
	elements []T
}

func (q *Queue) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() T {
	if len(q.elements) == 0 {
		return nil
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (q *Queue) PrintElements() {
	for _, v := range q.elements {
		fmt.Printf("%v\n", v)
	}
}
