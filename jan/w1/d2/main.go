package main

import (
	"fmt"
	"w1/d2/queue"
	"w1/d2/stack"
)

func main() {
	var st stack.Stack[int]
	var qq queue.Queue[int]

	st.Push(1)
	st.Peek()
	st.Push(3)
	st.Push(4)
	st.Push(3)
	st.Peek()

	qq.Enqueue(4)

	fmt.Println("stack", st)
	fmt.Println("queue", qq)
}
