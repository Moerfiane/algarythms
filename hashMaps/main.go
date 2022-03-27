package main

import (
	"fmt"
)

//array size is the size of the hash table array
const ArraySize = 7

//HashTable Structure
type HashTable struct {
	array [ArraySize]*bucket
}

//Bucket Structure
type bucket struct {
	head *bucketNode
}

//bucket node structure holding the linked list that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

//Functions for hastable
//Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

//Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//Functions for buckets
//Insert will take in a key, creake a node with the key and insert the node into the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Allready exists")
	}
}

//Search will take in a key and return true if the bucket has that value
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//Delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next

	}
}

//hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//init
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
func main() {
	hashTable := Init()
	list := []string{
		"Eric",
		"Kenny",
		"Kyle",
		"Stan",
		"Randy",
		"Butters",
		"Token",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("Stan")

	fmt.Println("Stan", hashTable.Search("Stan"))
	fmt.Println("Kenny", hashTable.Search("Kenny"))

}
