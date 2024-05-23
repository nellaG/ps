package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value int
	index int
	next  *Node
	prev  *Node
	poped bool
}

func (d *Deque) Push(x int) {
	newNode := &Node{value: x, poped: false}
	if d.head == nil {
		d.head = newNode
		d.last = d.head
		d.last.next = d.head
		d.last.prev = d.head
		d.head.index = 1

	} else {
		idx := d.last.index
		tail := d.last
		d.last.next = newNode
		newNode.prev = d.last
		d.last = newNode
		d.last.next = d.head
		d.last.prev = tail
		d.last.index = idx + 1
		d.head.prev = newNode

	}
}

type Deque struct {
	head *Node
	last *Node
}

func (d *Deque) Remove(node *Node) {
	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
	node = nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, x, idx int
	fmt.Fscanln(reader, &n)
	d := Deque{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		d.Push(x)
	}
	current := d.head
	for j := 0; j < n; j++ {
		val := current.value
		if current.index == 1 {

			d.Remove(current)
			fmt.Fprintf(writer, "%d ", 1)
			j++
		}
		if val > 0 {
			for i := 0; i < val; i++ {
				current = current.next
			}
		} else {
			for i := 0; i > val; i-- {
				current = current.prev

			}
		}
		idx = current.index
		d.Remove(current)
		fmt.Fprintf(writer, "%d ", idx)
	}
	writer.Flush()
}
