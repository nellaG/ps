package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscanln(reader, &a, &b)

	fmt.Printf("%d\n%d\n%d\n%d\n%d\n", a+b, a-b, a*b, a/b, a%b)
}
