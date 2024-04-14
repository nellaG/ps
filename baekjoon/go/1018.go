package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getCount(n, m int, c [][]string) int {
	odd := [][]string{
		{"W", "B", "W", "B", "W", "B", "W", "B"},
		{"B", "W", "B", "W", "B", "W", "B", "W"},
		{"W", "B", "W", "B", "W", "B", "W", "B"},
		{"B", "W", "B", "W", "B", "W", "B", "W"},
		{"W", "B", "W", "B", "W", "B", "W", "B"},
		{"B", "W", "B", "W", "B", "W", "B", "W"},
		{"W", "B", "W", "B", "W", "B", "W", "B"},
		{"B", "W", "B", "W", "B", "W", "B", "W"},
	}
	even := [][]string{
		odd[1], odd[0],
		odd[1], odd[0],
		odd[1], odd[0],
		odd[1], odd[0],
	}
	var countOdd, countEven int
	for i := n; i < n+8; i++ {
		for j := m; j < m+8; j++ {
			if c[i][j] != odd[i-n][j-m] {
				countOdd++
			}
			if c[i][j] != even[i-n][j-m] {
				countEven++
			}
		}
	}
	return min(countOdd, countEven)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, count, min_ int
	fmt.Fscanln(reader, &n, &m)
	chess := make([][]string, n)
	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		row := make([]string, m)
		chess[i] = row
		for j, v := range input {
			chess[i][j] = string(v)
		}
	}
	min_ = 64
	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			count = getCount(i, j, chess)
			if min_ > count {
				min_ = count
			}
		}
	}
	fmt.Println(min_)
}
