package trie

import (
	"errors"
	"fmt"
	"reflect"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	node       map[string]*TrieNode
	endOfWord  bool
	atrributes map[string]interface{}
}

func CreateTrie() *Trie {
	return &Trie{root: createEmptyNode()}
}

func createEmptyNode() *TrieNode {
	newNode := new(TrieNode)
	newNode.node = make(map[string]*TrieNode)
	newNode.atrributes = make(map[string]interface{})
	newNode.endOfWord = false
	return newNode
}

func (n *TrieNode) AddAttribute(attributeName string, attributeValue interface{}) *TrieNode {

	n.atrributes[attributeName] = attributeValue
	return n
}

func (n *TrieNode) GetAttributeValue(attributeName string) interface{} {
	return n.atrributes[attributeName]
}

func (n *TrieNode) GetValues() []interface{} {
	values := make([]interface{}, 0, len(n.atrributes))
	for _, v := range n.atrributes {
		values = append(values, v)
	}
	return values
}

func (n *TrieNode) GetKeys() []string {
	keys := make([]string, 0, len(n.atrributes))

	for k, _ := range n.atrributes {
		keys = append(keys, k)
	}
	return keys
}

func (n *Trie) Insert(key string) *TrieNode {
	var r rune
	var char string
	var isRootNode bool = true
	searchNode := n.root
	for _, r = range key {
		char = string(r)
		value, exist := searchNode.node[char]
		if exist {
			searchNode = value
			isRootNode = false
		} else {
			emptyNode := createEmptyNode()
			searchNode.node[char] = emptyNode
			searchNode = emptyNode
			isRootNode = false
		}
	}

	if !isRootNode {
		searchNode.endOfWord = true
	}
	return searchNode
}
func (n *Trie) GetNode(key string) *TrieNode {
	node, exist := n.getTrieNode(key)
	if !exist {
		return nil
	}

	return node
}

func (n *Trie) getTrieNode(key string) (*TrieNode, bool) {
	var (
		r     rune
		exist bool
		value *TrieNode
	)
	searchNode := n.root
	for _, r = range key {
		char := string(r)
		value, exist = searchNode.node[char]
		if exist {
			searchNode = value
		}
	}

	if exist {
		if !searchNode.endOfWord {
			return nil, false
		}
	}

	return searchNode, exist
}

func (n *Trie) Delete(key string) (bool, error) {

	type nodeTreeArray struct {
		searchNode *TrieNode
		char       string
	}

	var (
		r         rune
		exist     bool
		value     *TrieNode
		nodeArray = make([]nodeTreeArray, 0)
	)
	searchNode := n.root

	for _, r = range key {
		char := string(r)
		value, exist = searchNode.node[char]
		nodeArray = append(nodeArray, nodeTreeArray{value, char})
		if exist {
			searchNode = value
		}
	}

	if exist {
		if !searchNode.endOfWord {
			return false, errors.New("Key not found")
		} else {
			for i := len(nodeArray) - 1; i >= 0; i-- {
				deleteNode := nodeArray[i].searchNode

				j := i - 1
				if len(deleteNode.node) == 0 {
					if j >= 0 {
						delete(nodeArray[j].searchNode.node, nodeArray[i].char)
					} else {
						delete(n.root.node, nodeArray[i].char)
						n.root.endOfWord = true
					}
				} else if deleteNode.endOfWord {
					deleteNode.endOfWord = false
				}
			}
		}
	} else {
		return false, errors.New("Key not found")
	}

	return true, nil
}

func (T *Trie) Print() {
	fmt.Println("Root")
	fmt.Println("----------------------------------------------")
	T.root.Print()
	fmt.Println("----------------------------------------------")
}

func (tN *TrieNode) Print() {
	fmt.Println("TrieNode")
	fmt.Println("----------------------------------------------")
	fmt.Println("Attributes")
	fmt.Println("----------------------------------------------")
	for k, v := range tN.atrributes {
		fmt.Println("key : ", k)
		fmt.Println("value : ", reflect.ValueOf(v), " TypeOf :", reflect.TypeOf(v))
		tN, okN := v.(*TrieNode)
		if okN {
			tN.Print()
		} else {

			t, ok := v.(*Trie)
			if ok {
				t.Print()
			}
		}
	}
	fmt.Println("----------------------------------------------")
	for k, v := range tN.node {
		fmt.Println("key : ", k)
		fmt.Println("value : ", reflect.ValueOf(v), " TypeOf :", reflect.TypeOf(v))
		v.Print()
	}
}
