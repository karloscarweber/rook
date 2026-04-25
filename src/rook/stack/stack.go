// stack.go
//
// a stack is meant to be a byte slice with convenience methods
// to grow or shrink the slice, to append, to pop, and to remove
// elements from the slice. The slice is of bytes.

package stack

// import (
// 	"string"
// )

type Stackable interface {
	Pop()
	Push()
	Remove()
	Insert()
}

type Stack []byte

// pops a string from a byte slice and returns it.
func (stack Stack) Pop() byte {
	// str := []byte(stack)
	// last := len(str) - 1
	// // lastCh := str[last]
	bite := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return bite
}
