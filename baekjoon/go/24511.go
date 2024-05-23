package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, m int
	fmt.Fscan(reader, &n)
	d := list.New()
	type_ := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &type_[i])
	}
	for i := 0; i < n; i++ {
		var val int
		fmt.Fscan(reader, &val)
		if type_[i] == 0 {
			d.PushFront(val)
		}
	}
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var val int
		fmt.Fscan(reader, &val)
		d.PushBack(val)
		fmt.Fprintf(writer, "%d ", d.Front().Value)
		d.Remove(d.Front())
	}
	fmt.Fprintf(writer, "\n")
	writer.Flush()
}
