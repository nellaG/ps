package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

type member struct {
	age   int
	order int
	name  string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	defer writer.Flush()
	fmt.Fscanln(reader, &n)
	mbs := make([]member, n)
	for i := range mbs {
		fmt.Fscanln(reader, &mbs[i].age, &mbs[i].name)
		mbs[i].order = i
	}

	slices.SortStableFunc(mbs, func(a, b member) int {
		return cmp.Compare(a.age, b.age)
	})
	for _, k := range mbs {
		fmt.Fprintf(writer, "%d %s\n", k.age, k.name)
	}
}
