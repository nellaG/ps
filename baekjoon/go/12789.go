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

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pick() int {
	if len(*s) == 0 {
		return -1
	}
	index := len(*s) - 1
	return (*s)[index]

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x int
	c := 1
	fmt.Fscanln(reader, &n)
	s := Stack{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		if c == x {
			c++
		} else {
			s.Push(x)
		}
		for {
			if s.Pick() != c {
				break
			}
			s.Pop()
			c++
		}
	}
	if s.IsEmpty() {
		fmt.Println("Nice")
	} else {
		fmt.Println("Sad")
	}
}
