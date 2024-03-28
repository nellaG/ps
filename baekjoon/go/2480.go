package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	elements := strings.Fields(input)

	counts := make(map[int]int)
	revCounts := make(map[int]int)
	for _, elm := range elements {
		num, _ := strconv.Atoi(elm)
		counts[num]++
	}
	for key, value := range counts {
		revCounts[value] = key
	}
	keys := make([]int, len(revCounts))
	for k := range revCounts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })
	if keys[0] == 3 {
		fmt.Println(10000 + revCounts[3]*1000)
	} else if keys[0] == 2 {
		fmt.Println(1000 + revCounts[2]*100)
	} else {
		countKeys := make([]int, len(counts))
		for k := range counts {
			countKeys = append(countKeys, k)
		}
		fmt.Println(slices.Max(countKeys) * 100)
	}

}
