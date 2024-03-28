package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a int
	fmt.Fscanln(reader, &a)
	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}
}
