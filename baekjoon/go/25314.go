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
	for i := 0; i < a/4; i++ {
		fmt.Printf("%s ", "long")
	}
	fmt.Printf("%s\n", "int")
}
