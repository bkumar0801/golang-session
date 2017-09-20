package ds

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestEmptyDoublyLinkedList(t *testing.T) {
	node := Node{}
	assert.Equal(t, node.IsEmpty(), true)
}

func TestDoublyLinkedListWithOneNode(t *testing.T) {
	node := Node{}
	node.Add("test")
	assert.NotEqual(t, Head, nil)
	assert.NotEqual(t, Tail, nil)
	assert.Equal(t, Head, Tail)
	assert.Equal(t, Head.prev, (*Node)(nil))
	assert.Equal(t, Head.next, (*Node)(nil))
	assert.Equal(t, Tail.prev, (*Node)(nil))
	assert.Equal(t, Tail.next, (*Node)(nil))
}

func TestDoublyLinkeListWithMoreThanOneNode(t *testing.T) {
	node := Node{}
	node.Add("test")
	node.Add("success")
	assert.NotEqual(t, Head, nil)
	assert.NotEqual(t, Tail, nil)
	assert.NotEqual(t, Head, Tail)
	assert.Equal(t, Head.prev, (*Node)(nil))
	assert.Equal(t, Head.next, Tail)
	assert.Equal(t, Tail.prev, Head)
	assert.Equal(t, Tail.next, (*Node)(nil))
}


