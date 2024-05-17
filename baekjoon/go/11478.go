package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var s string
	fmt.Fscanln(reader, &s)
	length := len(s)
	map_ := make(map[string]int, 10000)
	for i := 0; i < length; i++ {
		for j := i; j <= length; j++ {
			map_[s[i:j]] = 1
		}
	}
	fmt.Fprintln(writer, len(map_)-1)
}
