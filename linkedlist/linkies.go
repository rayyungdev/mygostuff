package linkedlist

import (
	"fmt"
)

type valInterface interface {
	string | int
}
type Node[T valInterface] struct {
	data T
	next *Node[T]
}

func (n *Node[T]) SetData(data T) {
	n.data = data
}

func (n *Node[T]) Data() T {
	return n.data
}

func (n *Node[T]) NextNode() *Node[T] {
	return n.next
}

func (n *Node[T]) SetNext(node *Node[T]) {
	n.next = node
}

type LinkedList[T valInterface] struct {
	Head   *Node[T]
	Length int
}

func (l *LinkedList[T]) Preprend(n *Node[T]) {
	second := l.Head
	l.Head = n
	l.Head.next = second
	l.Length++
}

func (l *LinkedList[T]) DeleteWithValue(value T) {

	if l.Length == 0 {
		return
	}

	if l.Head.data == value {
		l.Head = l.Head.next
		l.Length--
		return
	}

	fmt.Println("Deleting Node with value: ", value)
	previousToDelete := l.Head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.Length--
}

func (l LinkedList[T]) PrintListData() {
	toPrint := l.Head
	for l.Length != 0 {
		fmt.Printf("%+v\n", toPrint.data)
		toPrint = toPrint.next
		l.Length--
	}
	fmt.Printf("\n")
}
