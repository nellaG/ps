package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, count int
	var left, bottom int
	fmt.Fscanln(reader, &n)

	defer writer.Flush()
	mat := make([][]int, 100)
	for j := 0; j < 100; j++ {
		mat[j] = make([]int, 100)
	}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &left, &bottom)
		for j := left; j < left+10; j++ {
			for k := bottom; k < bottom+10; k++ {
				if mat[j][k] != 1 {
					mat[j][k] = 1
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
