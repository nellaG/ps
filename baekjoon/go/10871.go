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
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var n, x int

	fmt.Fscanln(reader, &n, &x)
	input, _ := reader.ReadString('\n')
	elements := strings.Fields(input) // split
	for _, val := range elements {
		a, _ := strconv.Atoi(val)
		if a < x {
			fmt.Fprintf(writer, "%d ", a)
		}
	}
	fmt.Fprintf(writer, "\n")
	writer.Flush()
}
