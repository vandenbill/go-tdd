package queue

import "fmt"

var ErrNoSuchElement = fmt.Errorf("no such element")

type Node struct {
	Value int
	Next  *Node
}

func NewNode(v int) *Node {
	return &Node{Value: v}
}

type Queue struct {
	head   *Node
	Length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return q.Length == 0
}

func (q *Queue) Enqueue(node *Node) {
	second := q.head
	q.head = node
	q.head.Next = second
	q.Length += 1
}

func (q *Queue) Dequeue() (int, error) {
	result := 0
	if q.IsEmpty() {
		return 0, ErrNoSuchElement
	}

	if q.head.Next == nil {
		result = q.head.Value
		q.head = nil
		q.Length--
		return result, nil
	}

	length := q.Length
	head := q.head
	for length != 0 {
		if length == 2 {
			result = head.Next.Value
			head.Next = nil
			break
		}
		head = head.Next
		length--
	}
	q.Length--
	return result, nil
}
