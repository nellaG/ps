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

func (q *Queue) IsEmpty() int {
	if len(*q) != 0 {
		return 0
	}
	return 1
}
func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Front() int {
	if len(*q) == 0 {
		return -1
	}
	return (*q)[0]

}
func (q *Queue) Back() int {
	if len(*q) == 0 {
		return -1
	}
	index := len(*q) - 1
	return (*q)[index]

}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, x int
	var in string
	q := Queue{}
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &in)
		if in == "push" {
			fmt.Fscan(reader, &x)
			q.Push(x)
		} else {
			switch in {
			case "front":
				fmt.Fprintln(writer, q.Front())
			case "back":
				fmt.Fprintln(writer, q.Back())
			case "size":
				fmt.Fprintln(writer, q.Size())
			case "empty":
				fmt.Fprintln(writer, q.IsEmpty())
			case "pop":
				fmt.Fprintln(writer, q.Pop())
			}
		}
	}
	writer.Flush()
}
