package _622_design_circular_queue

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMyCircularQueue(t *testing.T) {
	queue := Constructor(3)
	assert.Check(t, is.Equal(queue.Rear(), -1))
	assert.Check(t, is.Equal(queue.Front(), -1))
	assert.Check(t, is.Equal(queue.EnQueue(1), true))
	assert.Check(t, is.Equal(queue.EnQueue(2), true))
	assert.Check(t, is.Equal(queue.EnQueue(3), true))
	assert.Check(t, is.Equal(queue.EnQueue(4), false))
	assert.Check(t, is.Equal(queue.Rear(), 3))
	assert.Check(t, is.Equal(queue.IsFull(), true))
	assert.Check(t, is.Equal(queue.DeQueue(), true))
	assert.Check(t, is.Equal(queue.IsFull(), false))
	assert.Check(t, is.Equal(queue.EnQueue(4), true))
	assert.Check(t, is.Equal(queue.Rear(), 4))
}
