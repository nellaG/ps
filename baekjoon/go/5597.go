package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	stud = 28
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num int
	check := make([]int, 30)

	for i := 0; i < stud; i++ {
		fmt.Fscanln(reader, &num)
		check[num-1] = 1
	}

	for i := 0; i < 30; i++ {
		if check[i] != 1 {
			fmt.Println(i + 1)
		}
	}
}
