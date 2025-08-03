package stack

type Stack struct {
	stack []int
	n     int
}

func NewStack() Stack {
	return Stack{
		stack: make([]int, 0),
		n:     0,
	}
}

func (stack *Stack) Pop() any {
	if stack.n == 0 {
		return nil
	}
	old := stack.stack
	item := old[stack.n-1:]
	stack.stack = old[:stack.n-1]
	stack.n--
	return item[0]
}

func (stack *Stack) Peek() any {
	if stack.n == 0 {
		return nil
	}
	return stack.stack[stack.n-1]
}

func (stack *Stack) Push(item int) {
	stack.stack = append(stack.stack, item)
	stack.n++
}
