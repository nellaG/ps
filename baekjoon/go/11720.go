package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num int
	var word string
	sum := 0
	fmt.Fscanln(reader, &num)
	fmt.Fscanln(reader, &word)
	for _, v := range word {
		val, _ := strconv.Atoi(string(v))
		sum += val
	}
	fmt.Println(sum)
}
