package doublelinkedlist

type Reader struct {
	list *DoubleLinkedList
}

func (reader *Reader) First() *Node {
	return reader.list.first
}

func (reader *Reader) Last() *Node {
	return reader.list.last
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Previous() *Node {
	return n.previous
}

func (list *DoubleLinkedList) GetReader() *Reader {
	reader := new(Reader)
	reader.list = list.deepCopy()
	return reader
}

func (list *DoubleLinkedList) Sync(reader *Reader) {
	reader.list = list.deepCopy()
}
