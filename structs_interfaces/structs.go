package structsinterfaces

import "fmt"

type stack struct {
	index int
	data  [5]int
}

func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

func Structs() {
	s := new(stack)
	s.push(23)
	s.push(14)
	s.push(22)
	s.push(15)
	fmt.Printf("the stack: %v\n", *s)
	// s.pop()
}
