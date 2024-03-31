package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var w string
	fmt.Scanln(&w)
	n := len(w)

	for i := 0; i < n/2; i++ {
		if w[i] != w[n-1-i] {
			fmt.Fprint(writer, 0)
			return
		}
	}
	fmt.Fprint(writer, 1)
}
