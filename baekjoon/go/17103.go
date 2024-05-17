package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, x int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscanln(reader, &x)
		var (
			arr [1000002]bool
		)
		for j := 2; j*j <= x; j++ {
			if arr[j] == false {
				for k := j + j; k <= x; k += j {
					arr[k] = true
				}
			}
		}
		for i := 2; i <= x/2; i++ {
			if arr[i] == false && arr[x-i] == false {
				c++
			}
		}
		fmt.Fprintln(writer, c)
	}
	writer.Flush()
}
