package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var day, night, v, ans int
	fmt.Fscanln(reader, &day, &night, &v)
	ans = (v - night) / (day - night)
	if (v-night)%(day-night) != 0 {
		ans++
	}
	fmt.Println(ans)
}
