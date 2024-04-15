package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input, _ := reader.ReadString('\n')
	num := make([]int, len(input))
	for i, v := range input {
		s, _ := strconv.Atoi(string(v))
		num[i] = s
	}
	sort.Ints(num)
	for i := len(num) - 1; i > 0; i-- {
		fmt.Fprint(writer, num[i])
	}
}
