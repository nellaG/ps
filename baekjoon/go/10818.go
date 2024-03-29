package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscanln(reader, &num)
	input, _ := reader.ReadString('\n')
	elements := strings.Fields(input) // split
	max_, _ := strconv.Atoi(elements[0])
	min_ := max_
	for _, v := range elements {
		val, _ := strconv.Atoi(v)
		fmt.Println(val)
		if val > max_ {
			max_ = val
		} else if val < min_ {
			min_ = val
		}
	}
	fmt.Printf("%d %d", min_, max_)
}
