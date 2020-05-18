package doublelinkedlist

import (
	"fmt"
	"reflect"
	"sync"
)

type Node struct {
	Key      []byte
	Value    interface{}
	previous *Node
	next     *Node
}

type DoubleLinkedList struct {
	first  *Node
	last   *Node
	count  int
	sorted bool
	backup *DoubleLinkedList
	mutex  sync.Mutex
}

func CreateDoubleLinkedList() *DoubleLinkedList {
	return new(DoubleLinkedList)
}

func CreateSortedDoubleLinkedList() *DoubleLinkedList {
	dbl := new(DoubleLinkedList)
	dbl.sorted = true
	return dbl
}

func createNode() *Node {
	return new(Node)
}
func setFirstAndLast(list *DoubleLinkedList, n *Node) {
	if n.previous == nil {
		list.first = n
	} else if n.next == nil {
		list.last = n
	}
}

func (list *DoubleLinkedList) Count() int {
	return list.count
}

func (list *DoubleLinkedList) IsSorted() bool {
	return list.sorted
}

func (list *DoubleLinkedList) Print() {
	node := list.first
	keys := make([]string, 0)
	values := make([]interface{}, 0)
	for node != nil {
		keys = append(keys, string(node.Key))
		values = append(values, reflect.ValueOf(node.Value))
		node = node.next
	}
	fmt.Println("Keys : ", keys)
	fmt.Println("Values : ", values)
}

func (list *DoubleLinkedList) deepCopy() *DoubleLinkedList {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	bkplist := new(DoubleLinkedList)
	if list.first == nil && list.last == nil {
		bkplist.first = nil
		bkplist.last = nil
		return bkplist
	}

	var copynode, currentnode, copynextnode *Node

	currentnode = list.first
	copynode = new(Node)
	bkplist.first = copynode
	bkplist.sorted = list.sorted
	bkplist.count = list.count

	for currentnode != nil {
		copynode.Key = currentnode.Key
		copynode.Value = currentnode.Value

		copynextnode = new(Node)
		copynode.next = copynextnode
		copynextnode.previous = copynode
		currentnode = currentnode.next
		copynode = copynode.next
	}
	if copynode.previous != nil {
		copynode.previous.next = nil
	}
	bkplist.last = copynode.previous
	return bkplist
}
