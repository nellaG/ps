package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) Pop() int {
	if len(*q) == 0 {
		return -1
	}
	pop := (*q)[0]
	*q = (*q)[1:]
	return pop
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	q := Queue{}
	fmt.Fscanln(reader, &n, &k)
	ans := make([]string, n)
	arr := make([]int, k-1)
	for i := 1; i <= n; i++ {
		q.Push(i)
	}
	for i := 0; i < n; i++ {
		len_ := len(q)
		tr := k % len_
		if k%len_ == 0 {
			tr = len_
		}
		for j := 0; j < tr-1; j++ {
			arr[j] = q.Pop()
		}
		ans[i] = strconv.Itoa(q.Pop())
		for j := 0; j < tr-1; j++ {
			q.Push(arr[j])
		}
	}
	fmt.Printf("<%s>\n", strings.Join(ans, ", "))
}
