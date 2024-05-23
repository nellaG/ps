package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StackStr []string

func (s *StackStr) PushStr(x string) {
	*s = append(*s, x)
}

func (s *StackStr) Pop() {
	index := len(*s) - 1
	*s = (*s)[:index]
}
func (s *StackStr) PickStr() string {
	if len(*s) == 0 {
		return ""
	}
	index := len(*s) - 1
	return (*s)[index]

}
func (s *StackStr) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	var input string
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		s := StackStr{}
		fmt.Fscanln(reader, &input)
		elems := strings.Split(input, "")
		for _, x := range elems {
			if x == "(" {
				s.PushStr(x)
			} else {
				if s.PickStr() == "(" {
					s.Pop()
				} else {
					s.PushStr(x)
				}
			}
		}
		if !s.IsEmpty() {
			fmt.Fprintln(writer, "NO")
		} else {
			fmt.Fprintln(writer, "YES")
		}
	}
	writer.Flush()
}
