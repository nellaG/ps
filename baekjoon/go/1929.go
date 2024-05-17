package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPrime(i int) bool {
	if i == 1 {
		return false
	}
	for j := 2; j*j <= i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var x, y int
	fmt.Fscanln(reader, &x, &y)
	arr := make([]int, y+1)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i*i <= y; i++ {
		if arr[i] == 0 {
			for j := i + i; j <= y; j += i {
				arr[j] = 1
			}
		}
	}
	for i := x; i <= y; i++ {
		if arr[i] == 0 {
			fmt.Println(i)
		}
	}
}
