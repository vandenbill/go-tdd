package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// - an queue has 0 length on construction
// - an queue is empty on construction
// - after n prepend to an empty queue, the queue is non empty and its length equal n
// - after we prepend a value, the value must be in frist order
// - if we take a data fromq queue, the data must be in the last order

func TestQueue(t *testing.T) {
	t.Run("an queue has 0 length on construction", func(t *testing.T) {
		q := queue.New()
		assert.True(t, q.isEmpty())
	})
}
