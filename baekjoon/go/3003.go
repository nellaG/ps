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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	ans := []int{1, 1, 2, 2, 2, 8}
	input, _ := reader.ReadString('\n')
	w := strings.Fields(input)

	for i := 0; i < 6; i++ {
		val, _ := strconv.Atoi(w[i])
		fmt.Fprintf(writer, "%d ", ans[i]-val)
	}
}
