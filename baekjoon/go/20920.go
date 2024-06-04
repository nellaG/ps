package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type word struct {
	str   string
	count int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	var s string
	fmt.Fscanln(reader, &n, &m)
	map_ := make(map[string]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		if len(s) < m {
			continue
		}
		map_[s] += 1
	}
	arr := make([]word, len(map_))
	i := 0
	for k, v := range map_ {
		arr[i] = word{str: k, count: v}
		i++
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].count != arr[j].count {
			return arr[i].count > arr[j].count
		}
		if len(arr[i].str) != len(arr[j].str) {
			return len(arr[i].str) > len(arr[j].str)
		}
		return arr[i].str < arr[j].str
	})
	for _, v := range arr {
		fmt.Fprintln(writer, v.str)
	}

}
