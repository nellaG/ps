package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, c, sum int
	var s string
	fmt.Fscanln(reader, &n)
	var map_ = make(map[string]bool)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		if s == "ENTER" {
			sum += c
			c = 0
			map_ = make(map[string]bool)
		} else {
			if !map_[s] {
				c++
				map_[s] = true
			}
		}
	}
	sum += c
	fmt.Println(sum)
}
