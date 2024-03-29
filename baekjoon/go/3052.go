package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	div = 42
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num int
	var count int
	check := make([]int, 42)

	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &num)
		check[num % div]++
	}

	for i := 0; i < div; i++ {
		if check[i] != 0 {
			count++
		}
	}
	fmt.Println(count)
}
