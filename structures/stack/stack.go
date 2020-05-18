package stack

import (
	"fmt"
	"github.com/jeyaprabhuj/forty/structures/linkedlist/doublelinkedlist"
	"reflect"
	"sync"
)

type Stack struct {
	list  *doublelinkedlist.DoubleLinkedList
	mutex sync.Mutex
}

func CreateNewStack() *Stack {
	s := new(Stack)
	s.list = doublelinkedlist.CreateDoubleLinkedList()
	return s
}

func (s *Stack) Push(key []byte, value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.list.Insert(key, value)

}

func (s *Stack) Pop() ([]byte, interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	last := s.list.GetReader().Last()
	if last != nil {
		key := last.Key
		value := last.Value
		s.list.DeleteLast()
		return key, value
	}
	return nil, nil
}

func (s *Stack) IsEmpty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.list.Count() > 0 {
		return false
	}

	return true
}

func (s *Stack) Print() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	last := s.list.GetReader().Last()
	if last != nil {
		key := last.Key
		value := last.Value

		fmt.Println(string(key), reflect.ValueOf(value))
	}

}
