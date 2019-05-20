package main

import (
	"fmt"
	"github.com/Luna-CY/go-learn/data_structs"
	"log"
)

func main() {
	str := "1 2 # 3 # # 4 # #"

	tree := &data_structs.BTreeNode{}
	err := data_structs.CreateNewBTree(tree, str)

	if nil != err {
		log.Fatalln(err)
	}

	fmt.Println(tree.PreOrderTraverse())
	fmt.Println(tree.InOrderTraverse())
	fmt.Println(tree.PostOrderTraverse())
}
