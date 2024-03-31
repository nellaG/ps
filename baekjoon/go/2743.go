package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var word string
	fmt.Fscanln(reader, &word)
	fmt.Print(len(word))
}
