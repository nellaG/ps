package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var count int
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		fmt.Println(0)
		return
	}
	for _, v := range word {
		if string(v) == " " {
			count++
		}
	}
	fmt.Println(count + 1)
}
