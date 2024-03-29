package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num, count int
	var target string
	fmt.Fscanln(reader, &num)
	input, _ := reader.ReadString('\n')
	elements := strings.Fields(input) // split

	fmt.Fscanln(reader, &target)

	for _, val := range elements {
		if val == target {
			count++
		}
	}
	fmt.Println(count)
}
