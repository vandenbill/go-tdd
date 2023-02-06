package queue_test

import (
	"testing"

	queue "github.com/vandenbill/go-tdd"

	"github.com/stretchr/testify/assert"
)

// - an queue has 0 length on construction
// - an queue is empty on construction
// - if we dequeue an empty queue, we got an error: ErrNoSuchElement
// - after n enqueue to an empty queue, the queue is non empty and its length equal n
// - if we dequeue from queue, the data must be in the last order

func TestQueue(t *testing.T) {
	q := queue.NewQueue()

	t.Run("an queue has 0 length on construction", func(t *testing.T) {
		assert.Equal(t, 0, q.Length)
	})

	t.Run("an queue is empty on construction", func(t *testing.T) {
		assert.True(t, q.IsEmpty())
	})

	t.Run("if we dequeue an empty queue, we got an error: ErrNoSuchElement", func(t *testing.T) {
		_, err := q.Dequeue()
		assert.Equal(t, queue.ErrNoSuchElement, err)
	})

	t.Run("test queue that only have one node", func(t *testing.T) {
		n := queue.NewNode(9823783)
		q.Enqueue(n)

		v, err := q.Dequeue()
		assert.Nil(t, err)
		assert.Equal(t, n.Value, v)
	})

	n1 := queue.NewNode(1)
	n3 := queue.NewNode(3)
	n2 := queue.NewNode(2)
	n4 := queue.NewNode(4)
	n5 := queue.NewNode(5)
	q.Enqueue(n1)
	q.Enqueue(n2)
	q.Enqueue(n3)
	q.Enqueue(n4)
	q.Enqueue(n5)

	t.Run("after n enqueue to an empty queue, the queue is non empty and its length equal n", func(t *testing.T) {
		assert.Equal(t, 5, q.Length)
	})

	t.Run("if we dequeue from queue, the data must be in the last order", func(t *testing.T) {
		v, err := q.Dequeue()
		assert.Nil(t, err)
		assert.Equal(t, n1.Value, v)
	})
}
