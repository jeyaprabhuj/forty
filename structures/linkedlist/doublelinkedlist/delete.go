package doublelinkedlist

import (
	"bytes"
)

func (dbl *DoubleLinkedList) DeleteLast() {
	deleteList := dbl.deepCopy()
	deleteDblNode := deleteList.last

	if deleteDblNode != nil {
		if deleteDblNode.previous != nil {
			deleteList.last = deleteDblNode.previous
			deleteDblNode.previous.next = nil
			deleteDblNode.previous = nil

			deleteList.backup = deleteList.deepCopy()
		} else {
			deleteList.first = nil
			deleteList.last = nil
			deleteList.backup = nil
		}
		deleteList.count--
	}
	dbl.mutex.Lock()
	defer dbl.mutex.Unlock()
	dbl.first = deleteList.first
	dbl.last = deleteList.last
	dbl.backup = deleteList.backup
	dbl.count = deleteList.count
}

func (dbl *DoubleLinkedList) DeleteFirst() {
	deleteList := dbl.deepCopy()
	deleteDblNode := deleteList.first

	if deleteDblNode != nil {
		if deleteDblNode.next != nil {
			deleteList.first = deleteDblNode.next
			deleteDblNode.next.previous = nil
			deleteDblNode.next = nil

			deleteList.backup = deleteList.deepCopy()
		} else {
			deleteList.first = nil
			deleteList.last = nil
			deleteList.backup = nil
		}
		deleteList.count--
	}
	dbl.mutex.Lock()
	defer dbl.mutex.Unlock()
	dbl.first = deleteList.first
	dbl.last = deleteList.last
	dbl.backup = deleteList.backup
	dbl.count = deleteList.count
}

func (dbl *DoubleLinkedList) DeleteByKey(key []byte) {
	deleteList := dbl.deepCopy()
	searchnode := deleteList.first

	for searchnode != nil {
		if bytes.Compare(key, searchnode.Key) == 0 {
			if searchnode.previous != nil {
				searchnode.previous.next = searchnode.next
			}
			if searchnode.next != nil {
				searchnode.next.previous = searchnode.previous
			}
			deleteList.count--
			deleteList.backup = deleteList.deepCopy()
		} else {
			searchnode = searchnode.next
		}
	}
	dbl.mutex.Lock()
	defer dbl.mutex.Unlock()
	dbl.first = deleteList.first
	dbl.last = deleteList.last
	dbl.backup = deleteList.backup
	dbl.count = deleteList.count
}
