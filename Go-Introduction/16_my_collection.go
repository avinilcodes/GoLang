package main

//Why this is not working, is it because Go does not suppot pointer arithmetic
import "fmt"

type myCollection string
type Stack interface {
	Push()
	Pop() int
	Top() int
}

type node struct {
	data int
	next *node
}
type stack struct {
	n *node
}

func newNode(data int) *node {
	n := new(node)
	n.data = data
	n.next = nil
	return n
}
func (mycol myCollection) Push(data int, st stack) *stack {
	if st.n == nil {
		st.n = newNode(data)
		return &st
	}
	var t *node
	for t = st.n; t != nil; t = t.next {
	}
	t.next = newNode(data)
	return &st
}

func (mycol myCollection) Pop(st stack) *stack {
	t := st.n.data
	st.n = st.n.next
	fmt.Println("Top of stack :", t)
	return &st
}

func (mycol myCollection) Top(st stack) int {
	t := st.n.data
	return t
}

func printStack(st stack) {
	var t *node
	for t = st.n; t != nil; t = t.next {
		fmt.Println(t.data)
	}
}

func main() {
	stack := new(stack)

	var mycol myCollection
	stack = mycol.Push(1, *stack)
	stack = mycol.Push(2, *stack)
	stack = mycol.Push(3, *stack)
	stack = mycol.Push(4, *stack)
	stack = mycol.Push(5, *stack)
	printStack(*stack)
}
