package main

import (
	"fmt"
	"github.com/Luna-CY/go-learn/algorithm"
)

func main() {
	bst := algorithm.BSTree{}

	bst.Insert(124)
	bst.Insert(123)
	bst.Insert(80)
	bst.Insert(30)
	bst.Insert(60)
	bst.Insert(100)
	bst.Insert(120)
	bst.Insert(150)
	bst.Insert(130)

	fmt.Println(bst.PreOrderTraverse())
	fmt.Println(bst.InOrderTraverse())
}
