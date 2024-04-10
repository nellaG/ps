package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, count int
	fmt.Fscanln(reader, &n)
	input, _ := reader.ReadString('\n')
	for _, s := range strings.Fields(input) {
		ss, _ := strconv.Atoi(s)
		if ss == 1 {
			count++
			continue
		}
		for j := 2; j < ss; j++ {
			if ss%j == 0 {
				count++
				break
			}
		}
	}
	fmt.Println(n - count)
}
