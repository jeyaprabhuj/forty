package main

import (
	"fmt"
	"github.com/jeyaprabhuj/forty/structures/trie"
)

func main() {

	path := trie.CreateTrie()

	path.Insert("abc").AddAttribute("POST", "/api/v1/user/add")

	path.Insert("abe")
	path.Insert("abef")

	url, _ := path.GetNode("abc").GetAttributeValue("POST").(string)
	fmt.Println(url)

	path.Delete("abe")
	path.Delete("abef")
	path.Delete("abc")
	fmt.Println("Completed")

}
