package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var name, sign string
	fmt.Fscanln(reader, &n)
	strs := make(map[string]bool, n)
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &name, &sign)
		if sign == "enter" {
			strs[name] = true
		} else {
			strs[name] = false
		}
	}
	for k, v := range strs {
		if v {
			ss = append(ss, k)
		}
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] > ss[j]
	})
	for _, s := range ss {
		fmt.Fprintln(writer, s)
		if s == "" {
			break
		}
	}
}
