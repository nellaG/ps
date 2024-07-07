package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for {
		_, err := fmt.Fscanln(reader, &n)
		if err == io.EOF {
			return
		}
		fmt.Fprintln(writer, cantor(n))
	}
}

func cantor(n int) string {
	if n <= 0 {
		return "-"
	}
	return cantor(n-1) + strings.Repeat(" ", int(math.Pow(3, float64(n-1)))) + cantor(n-1)
}
