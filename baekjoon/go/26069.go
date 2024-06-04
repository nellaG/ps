package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s1, s2 string
	fmt.Fscanln(reader, &n)
	var map_ = make(map[string]bool)
	map_["ChongChong"] = true
	sum := 1
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s1, &s2)
		if map_[s1] != map_[s2] {
			sum++
			map_[s1] = true
			map_[s2] = true
		}
	}
	fmt.Println(sum)
}
