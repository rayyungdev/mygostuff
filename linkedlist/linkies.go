package linkedlist

import (
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

func (n *Node[T]) SetData(data T) {
	n.data = data
}

func (n *Node[T]) NextNode() *Node[T] {
	return n.next
}

type LinkedList[T comparable] struct {
	head   *Node[T]
	length int
}

func (l *LinkedList[T]) Preprend(n *Node[T]) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList[T]) DeleteWithValue(value T) {

	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	fmt.Println("Deleting Node with value: ", value)
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l LinkedList[T]) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%+v\n", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
