package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	index := len(*s) - 1
	pop := (*s)[index]
	*s = (*s)[:index]
	return pop
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x, c int
	fmt.Fscanln(reader, &n)
	s := Stack{}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x)
		if x == 0 {
			s.Pop()
		} else {
			s.Push(x)
		}
	}
	for _, x := range s {
		c += x
	}
	fmt.Println(c)
}
