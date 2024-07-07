package main

import (
	"bufio"
	"fmt"
	"os"
)

var c = 0
var ans = 0

func merge(arr []int, tmp []int, p, q, r, s int) {
	i := p
	j := q + 1
	t := 0
	for i <= q && j <= r {
		if arr[i] <= arr[j] {
			tmp[t] = arr[i]
			t++
			i++
		} else {
			tmp[t] = arr[j]
			t++
			j++
		}
	}
	for i <= q {
		tmp[t] = arr[i]
		t++
		i++
	}
	for j <= r {
		tmp[t] = arr[j]
		t++
		j++
	}
	i = p
	t = 0
	for i <= r {
		arr[i] = tmp[t]
		c++
		if c == s {
			ans = tmp[t]
		}
		t++
		i++
	}
}

func mergeSort(arr []int, tmp []int, p, r, s int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(arr, tmp, p, q, s)
		mergeSort(arr, tmp, q+1, r, s)
		merge(arr, tmp, p, q, r, s)
	}
}

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)
	arr := make([]int, n)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])

	}
	mergeSort(arr, tmp, 0, n-1, m)
	if c < m {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
