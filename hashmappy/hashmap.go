package hashmappy

import (
	"fmt"

	"github.com/rayyungdev/mygostuff/linkedlist"
)

const ArraySize = 7

// Hash Table structure
type HashTable[T string] struct {
	array [ArraySize]*Bucket[T]
}

func (h *HashTable[T]) Insert(key T) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable[T]) Search(key T) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable[T]) Delete(key T) {
	index := hash(key)
	h.array[index].delete(key)
}

func hash[T string](key T) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Bucket Structure -- LinkedList (Can be borrowed from my other repo)
type Bucket[T string] struct {
	linkedlist.LinkedList[T]
}

func (b *Bucket[T]) insert(k T) error {
	if b.search(k) {
		return fmt.Errorf("%+v already exists", k)
	}
	newNode := &linkedlist.Node[T]{}
	newNode.SetData(k)
	newNode.SetNext(b.Head)
	b.Head = newNode
	return nil
}

func (b *Bucket[T]) search(k T) bool {
	currentNode := b.Head
	for currentNode != nil {
		if currentNode.Data() == k {
			return true
		}
		currentNode = currentNode.NextNode()
	}
	return false
}

func (b *Bucket[T]) delete(k T) error {
	type exists struct {
		exist bool
	}

	doesthisexist := &exists{exist: false}

	if b.Head.Data() == k {
		b.Head = b.Head.NextNode()
		return nil
	}

	previousNode := b.Head
	for previousNode.NextNode() != nil {
		if previousNode.NextNode().Data() == k {
			previousNode.SetNext(previousNode.NextNode().NextNode())
			doesthisexist.exist = true
		}
		previousNode = previousNode.NextNode()
	}
	if !doesthisexist.exist {
		return fmt.Errorf("Key %+v not found", k)
	}
	return nil
}

func Init[T string]() *HashTable[T] {
	result := &HashTable[T]{}
	for i := range result.array {
		result.array[i] = &Bucket[T]{}
	}
	return result
}
