package doublelinkedlist

type snapshotOfLinkedList struct {
	list *DoubleLinkedList
}

func (snapshot *snapshotOfLinkedList) First() *Node {
	return snapshot.list.first
}
func (snapshot *snapshotOfLinkedList) Last() *Node {
	return snapshot.list.last
}

func (snapshot *snapshotOfLinkedList) Print() {
	snapshot.list.Print()
}

func (list *DoubleLinkedList) CreateSnapshot() *snapshotOfLinkedList {
	snaplist := new(snapshotOfLinkedList)
	snaplist.list = list.backup.deepCopy()
	return snaplist
}
