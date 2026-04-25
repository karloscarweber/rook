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
