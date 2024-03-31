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
	var n, m int
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscanln(reader, &n, &m)
	nmat := make([][]int, n)
	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		row := make([]int, m)
		for j, v := range strings.Fields(input) {
			row[j], _ = strconv.Atoi(v)
		}
		nmat[i] = row
	}
	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		for j, v := range strings.Fields(input) {
			w, _ := strconv.Atoi(v)
			nmat[i][j] += w
			fmt.Fprintf(writer, "%d ", nmat[i][j])
		}
		fmt.Fprintln(writer)
	}
}
