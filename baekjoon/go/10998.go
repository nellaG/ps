package main

import (
	"bufio"
	"fmt"
	"os"
)

func ps10998(a int, b int) int {
	return a * b
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscanln(reader, &a, &b)
	result := ps10998(a, b)
	fmt.Println(result)
}
