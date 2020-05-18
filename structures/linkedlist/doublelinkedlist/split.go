package doublelinkedlist

import (
	"bytes"
)

func (list *DoubleLinkedList) SplitOnIndex(index int) (*DoubleLinkedList, *DoubleLinkedList) {
	var currentnode *Node
	var currentindex int
	var list1, list2 *DoubleLinkedList

	currentindex = 0
	operatinglist := list.backup.deepCopy()
	currentnode = operatinglist.first

	list1 = new(DoubleLinkedList)
	list2 = new(DoubleLinkedList)

	if index-1 >= operatinglist.count {
		panic("Index out of range")
	}

	var splitnode1, splitnode2 *Node
	splitnode1 = new(Node)
	for currentnode != nil {
		//assign nodes to list1 until index is reached
		if currentindex <= index {
			if currentindex == 0 {
				list1.first = splitnode1
			}
			list1.count++
		} else {
			list2.count++
		}

		splitnode1.Key = currentnode.Key
		splitnode1.Value = currentnode.Value
		splitnode2 = new(Node)
		splitnode1.next = splitnode2
		splitnode2.previous = splitnode1

		currentnode = currentnode.next

		currentindex++
		if list2.count == 0 && currentindex > index {
			list1.last = splitnode1
			if splitnode1.next != nil {
				splitnode1.next = nil
			}
			splitnode1 = new(Node)
			list2.first = splitnode1
		} else {
			splitnode1 = splitnode1.next
		}

	}

	if splitnode1.previous != nil && list2.first != nil {
		splitnode1.previous.next = nil
		list2.last = splitnode1.previous
	} else if list2.count == 0 {
		list2.first = nil
		list2.last = nil
	}

	list1.sorted = operatinglist.sorted
	list1.backup = list1.deepCopy()
	list2.sorted = operatinglist.sorted
	list2.backup = list2.deepCopy()
	return list1, list2
}

func (list *DoubleLinkedList) SplitByKey(key []byte) (*DoubleLinkedList, *DoubleLinkedList) {
	var currentnode *Node
	var list1, list2, currentlist *DoubleLinkedList

	operatinglist := list.backup.deepCopy()
	currentnode = operatinglist.first

	list1 = new(DoubleLinkedList)
	list2 = new(DoubleLinkedList)
	currentlist = list1

	var splitnode1, splitnode2 *Node
	splitnode1 = new(Node)
	if currentnode != nil {
		currentlist.first = splitnode1
	}
	for currentnode != nil {
		if bytes.Compare(currentnode.Key, key) == 0 {
			list1.last = splitnode1
			if splitnode1.next != nil {
				splitnode1.next = nil
			}
			splitnode1 = new(Node)
			list2.first = splitnode1
			currentlist = list2
		}

		splitnode1.Key = currentnode.Key
		splitnode1.Value = currentnode.Value
		splitnode2 = new(Node)
		splitnode1.next = splitnode2
		splitnode2.previous = splitnode1
		currentnode = currentnode.next
		splitnode1 = splitnode1.next
		currentlist.count++
	}

	list1.sorted = operatinglist.sorted
	list1.backup = list1.deepCopy()
	list2.sorted = operatinglist.sorted
	list2.backup = list2.deepCopy()
	return list1, list2
}
