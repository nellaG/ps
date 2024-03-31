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
	max := []int{0, 0, 0}
	defer writer.Flush()
	nmat := make([][]int, 9)
	for i := 0; i < 9; i++ {
		input, _ := reader.ReadString('\n')
		row := make([]int, 9)
		for j, v := range strings.Fields(input) {
			vv, _ := strconv.Atoi(v)
			if max[2] < vv {
				max[0] = i
				max[1] = j
				max[2] = vv
			}
		}
		nmat[i] = row
	}
	fmt.Println(max[2])
	fmt.Printf("%d %d\n", max[0]+1, max[1]+1)
}
