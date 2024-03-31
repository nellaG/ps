package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num int
	var word string
	fmt.Fscanln(reader, &num)
	for i := 0; i < num; i++ {
		fmt.Fscanln(reader, &word)
		fmt.Printf("%c%c", word[0], word[len(word)-1])

	}
}
