package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var x1, y1, x2, y2, x3, y3 int
	fmt.Scanf("%d %d\n%d %d\n%d %d", &x1, &y1, &x2, &y2, &x3, &y3)
	if x1 == x2 {
		fmt.Fprintf(writer, "%d ", x3)
	} else if x1 == x3 {
		fmt.Fprintf(writer, "%d ", x2)
	} else {
		fmt.Fprintf(writer, "%d ", x1)
	}
	if y1 == y2 {
		fmt.Fprintln(writer, y3)
	} else if y1 == y3 {
		fmt.Fprintln(writer, y2)
	} else {
		fmt.Fprintln(writer, y1)
	}

}
