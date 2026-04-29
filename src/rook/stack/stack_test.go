// stack_test.go

package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	s := Stack{'a','b','c','d'}
	assert.Equal(t, 4, len(s), "Not enough tokens.")
	assert.Equal(t, s[0], byte('a'))
	b := s.Pop()
	assert.Equal(t, b, byte('d'))
}

func TestPush(t *testing.T) {
	s := Stack{'a','b','c','d'}
	assert.Equal(t, 4, len(s), "Not enough tokens.")
	s = s.Push('e')
	assert.Equal(t, 5, len(s), "Not enough tokens.")
	assert.Equal(t, s[4], byte('e'))
}

func TestRemove(t *testing.T) {
	s := Stack{'a','b','c','d'}
	assert.Equal(t, 4, len(s), "Not enough tokens.")
	s = s.Remove(1)
	assert.Equal(t, 3, len(s), "Not enough tokens.")
	assert.Equal(t, s[1], byte('c'))
}

func TestInsert(t *testing.T) {
	s := Stack{'a','b','c','d'}
	assert.Equal(t, 4, len(s), "Not enough tokens.")
	s = s.Insert(1, byte('w'))
	assert.Equal(t, 5, len(s), "Not enough tokens.")
	assert.Equal(t, s[1], byte('w'))
}
