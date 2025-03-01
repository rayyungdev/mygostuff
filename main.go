package main

import (
	"github.com/rayyungdev/mygostuff/linkedlist"
)

func linkedListExample() {
	/* Example from: https://www.youtube.com/watch?v=PGcTioRPBhU&t=212s */
	mylist := linkedlist.LinkedList[int]{}
	node1 := &linkedlist.Node[int]{}
	node1.SetData(48)
	node2 := &linkedlist.Node[int]{}
	node2.SetData(12)
	node3 := &linkedlist.Node[int]{}
	node3.SetData(18)
	node4 := &linkedlist.Node[int]{}
	node4.SetData(4)

	mylist.Preprend(node1)
	mylist.Preprend(node2)
	mylist.Preprend(node3)
	mylist.Preprend(node4)
	mylist.PrintListData()

	mylist.DeleteWithValue(12)

	mylist.PrintListData()

	emptyList := linkedlist.LinkedList[int]{}

	emptyList.DeleteWithValue(10)
	emptyList.PrintListData()
}

func main() {
	linkedListExample()
}
