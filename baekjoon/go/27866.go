package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var word string
	var n int
	fmt.Fscanln(reader, &word)
	fmt.Fscanln(reader, &n)

	fmt.Printf("%c", word[n-1])
}
