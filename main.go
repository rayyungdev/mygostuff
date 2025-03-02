package main

import (
	"fmt"

	"time"

	"github.com/rayyungdev/mygostuff/hashmappy"
	"github.com/rayyungdev/mygostuff/linkedlist"
	"github.com/rayyungdev/mygostuff/otherthings"
)

func LinkedListExample() {
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

func hashMapExample() {
	hashTable := hashmappy.Init[string]()
	list := []string{
		"ERIC",
		"KYLE",
		"STAN",
		"BUTTERS",
		"RANDY",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}
	result := hashTable.Search("KYLE")

	fmt.Println("FOUND KYLE: ", result)

	hashTable.Delete("KYLE")

	fmt.Println("Delete Kyle it shouldn't exist")
	fmt.Printf("Check Kyle: %+v", hashTable.Search("KYLE"))

}

func main() {
	client := otherthings.StartNewCollection(5*time.Minute, "Things")
	LinkedListExample()
	hashMapExample()
	client.AddToCollectionLocked("THINGS", map[string]string{"Hello": "No"})

	test, err := client.RetrieveItemLocked("THINGS")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%+v", *test)
}
