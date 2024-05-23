package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Deque struct {
	head   *Node
	length int
	last   *Node
}

func (d *Deque) Push(x int) {
	newNode := &Node{value: x}
	if d.head == nil {
		d.head = newNode
		d.length = 1
		d.last = d.head
	} else {
		d.last.next = newNode
		newNode.prev = d.last
		d.length += 1
		d.last = newNode
	}
}

func (d *Deque) PushLeft(x int) {
	newNode := &Node{value: x}
	if d.head == nil {
		d.head = newNode
		d.length = 1
		d.last = d.head
	} else {
		d.head.prev = newNode
		newNode.next = d.head
		d.head = newNode
		d.length += 1
	}
}

func (d *Deque) Pop() int {
	if d.head == nil {
		return -1
	}
	popNode := d.last
	d.last = popNode.prev
	d.length -= 1
	if d.length == 0 {
		d.head = nil
		d.last = nil
	}
	return popNode.value
}

func (d *Deque) PopLeft() int {
	if d.head == nil {
		return -1
	}
	pop := d.head
	d.head = d.head.next
	d.length -= 1
	if d.length == 0 {
		d.head = nil
		d.last = nil
	}
	return pop.value
}

func (d *Deque) Pick() int {
	if d.head == nil {
		return -1
	}
	return d.last.value
}
func (d *Deque) PickLeft() int {
	if d.head == nil {
		return -1
	}
	return d.head.value
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, x int
	fmt.Fscanln(reader, &n)
	d := Deque{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		switch x {
		case 1:
			{
				fmt.Fscan(reader, &x)
				d.PushLeft(x)
			}
		case 2:
			{
				fmt.Fscan(reader, &x)
				d.Push(x)
			}
		case 3:
			fmt.Fprintln(writer, d.PopLeft())
		case 4:
			fmt.Fprintln(writer, d.Pop())
		case 5:
			fmt.Fprintln(writer, d.length)
		case 6:
			{
				emp := 1
				if d.length != 0 {
					emp = 0
				}
				fmt.Fprintln(writer, emp)
			}
		case 7:
			fmt.Fprintln(writer, d.PickLeft())
		case 8:
			fmt.Fprintln(writer, d.Pick())
		}
	}
	writer.Flush()
}
