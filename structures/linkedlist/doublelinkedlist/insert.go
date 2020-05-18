package doublelinkedlist

import (
	"bytes"
)

func (list *DoubleLinkedList) Insert(Key []byte, Value interface{}) {

	list.mutex.Lock()

	var newnode *Node

	if list.first == nil {
		newnode = createNode()
		newnode.Key = Key
		list.first = newnode
		list.last = newnode
	} else {

		if list.sorted {
			newnode = sortedInsert(list, Key)
		} else {
			newnode = add(list, list.last, Key, true)
			setFirstAndLast(list, newnode)
		}
	}
	newnode.Value = Value
	list.count++
	list.mutex.Unlock()
	list.backup = list.deepCopy()

}

func sortedInsert(list *DoubleLinkedList, Key []byte) *Node {

	var newnode, previousNode, currentNode *Node
	currentNode = list.first

	for currentNode != nil {
		if bytes.Compare(Key, currentNode.Key) == -1 {
			newnode = add(list, currentNode, Key, false)
			setFirstAndLast(list, newnode)
			return newnode
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}

	newnode = add(list, previousNode, Key, true)
	setFirstAndLast(list, newnode)
	return newnode
}

func add(list *DoubleLinkedList, n *Node, Key []byte, append bool) *Node {
	newNode := createNode()
	newNode.Key = Key

	if append {

		if n.next != nil {
			if n.next.previous != nil {
				newNode.next = n.next
				n.next.previous = newNode

			}
		}
		n.next = newNode
		newNode.previous = n
	} else {
		//prepend
		newNode.next = n
		if n.previous != nil {
			if n.previous.next != nil {
				n.previous.next = newNode
			}
			newNode.previous = n.previous
		}

		n.previous = newNode

	}

	return newNode

}
