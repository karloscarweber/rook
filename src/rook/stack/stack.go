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
	bite := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return bite
}

func (stack Stack) Push(bt byte) Stack {
	stack = append(stack, bt)
	return stack
}

func (stack Stack) Remove(pl int) Stack  {
	copy(stack[pl:], stack[pl+1:])
	return stack[:len(stack)-1]
}

func (stack Stack) Insert(pl int, bt byte) Stack  {
	stack = append(stack, '0')
	copy(stack[pl+1:], stack[pl:])
	stack[pl] = bt
	return stack
}
