package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	word, _, _ := reader.ReadRune()
	fmt.Print(int(word))
}
