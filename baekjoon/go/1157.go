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
	var w string
	ab := make([]int, 26)
	smasc := 97
	fmt.Fscanln(reader, &w)

	for _, v := range w {
		if int(v) < smasc {
			v += 32
		}
		ab[v-97]++
	}
	var max = []int{0, ab[0]}
	for i, v := range ab {
		if v > max[1] {
			max = []int{i, v}
		}
	}
	for i, v := range ab {
		if max[0] != i && max[1] == v {
			fmt.Fprintln(writer, "?")
			return
		}
	}
	fmt.Fprintf(writer, "%c", max[0]+65)
}
