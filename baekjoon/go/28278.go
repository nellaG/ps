package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (s *Stack) Len() int {
	return len(*s)
}
func (s *Stack) IsEmpty() int {
	if s.Len() == 0 {
		return 1
	}
	return 0
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
	writer := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(reader, &n)
	s := Stack{}
	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		elems := strings.Fields(input)
		if elems[0] == "1" {
			val, _ := strconv.Atoi(elems[1])
			s.Push(val)
		} else {
			switch elems[0] {
			case "2":
				fmt.Fprintln(writer, s.Pop())
			case "3":
				fmt.Fprintln(writer, s.Len())
			case "4":
				fmt.Fprintln(writer, s.IsEmpty())
			case "5":
				fmt.Fprintln(writer, s.Pick())
			}
		}
	}
	writer.Flush()
}
