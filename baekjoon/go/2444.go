package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		fmt.Fprintln(writer, strings.Repeat(" ", n-i)+strings.Repeat("*", 2*i-1))
	}
	for i := n - 1; i > 0; i-- {
		fmt.Fprintln(writer, strings.Repeat(" ", n-i)+strings.Repeat("*", 2*i-1))
	}
}
