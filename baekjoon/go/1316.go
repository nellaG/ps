package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input, drop string
	var n, count int
	writer := bufio.NewWriter(os.Stdout)
	mem := make(map[string]int)
	defer writer.Flush()
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		for i := 97; i < 123; i++ {
			mem[string(i)] = -2
		}
		fmt.Fscanln(reader, &input)

		for i, v := range input {
			drop = ""
			w := string(v)
			if mem[w] == -2 {
				mem[w] = i
			} else {
				if i-mem[w] == 1 {
					mem[w] = i
				} else {
					drop += w
					break
				}
			}
		}
		if drop == "" {
			count++
		}
	}
	fmt.Println(count)
}
