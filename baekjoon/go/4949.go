package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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
	for {
		input, _ := reader.ReadString('\n')
		if input == ".\n" {
			break
		}
		s := StackStr{}
		for _, i := range input {
			x := string(i)
			runeX := []rune(x)[0]
			if unicode.IsLetter(runeX) || unicode.IsSpace(runeX) {
			} else {
				if x == "(" || x == "{" || x == "[" {
					s.PushStr(x)
				} else {
					switch x {
					case ")":
						if s.PickStr() == "(" {
							s.Pop()
						} else {
							s.PushStr(x)
						}
					case "]":
						if s.PickStr() == "[" {
							s.Pop()
						} else {
							s.PushStr(x)
						}
					case "}":
						if s.PickStr() == "{" {
							s.Pop()
						} else {
							s.PushStr(x)
						}
					}
				}
			}
		}
		if !s.IsEmpty() {
			fmt.Fprintln(writer, "no")
		} else {
			fmt.Fprintln(writer, "yes")
		}
		writer.Flush()
	}
}
