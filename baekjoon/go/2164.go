package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) Pop() int {
	if len(*q) == 0 {
		return -1
	}
	pop := (*q)[0]
	*q = (*q)[1:]
	return pop
}
func (q *Queue) Front() int {
	if len(*q) == 0 {
		return -1
	}
	return (*q)[0]

}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x int
	q := Queue{}
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		q.Push(i + 1)
	}
	for i := 0; i < n-1; i++ {
		q.Pop()
		x = q.Front()
		q.Pop()
		q.Push(x)
	}
	fmt.Println(q.Front())
}
