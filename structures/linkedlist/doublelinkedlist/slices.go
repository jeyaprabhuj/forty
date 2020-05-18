package doublelinkedlist

func (list *DoubleLinkedList) GetKeys() [][]byte {
	keys := make([][]byte, list.count)

	retrievelist := list.deepCopy()
	currentnode := retrievelist.first
	for currentnode != nil {
		keys = append(keys, currentnode.Key)
	}
	return keys
}

func (list *DoubleLinkedList) GetValues() []interface{} {
	values := make([]interface{}, list.count)

	retrievelist := list.deepCopy()
	currentnode := retrievelist.first
	for currentnode != nil {
		values = append(values, currentnode.Value)
	}
	return values
}
