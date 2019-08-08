package main

import (
	"fmt"
	"github.com/Luna-CY/go-learn/data_structs/list"
)

func main() {
	isl := list.IntSortList{}
	isl.Init(100, 10, 25, 2, 123, 21, 22, 5, 9, 101, 42, 31, 1)
	fmt.Println(isl.GetLength())
	fmt.Println(isl.BubbleSort())

	fmt.Println(isl.GetList())

	fmt.Println("-----------------------------")

	isl.Init(1, 2, 3, 4, 5, 6, 7, 8, 10, 9, 30, 20, 40)
	fmt.Println(isl.GetLength())
	fmt.Println(isl.BubbleSort())

	fmt.Println(isl.GetList())

	fmt.Println("-----------------------------")

	isl.Init(40, 30, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	fmt.Println(isl.GetLength())
	fmt.Println(isl.BubbleSort())

	fmt.Println(isl.GetList())

	fmt.Println("-----------------------------")

	isl.Init(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40)
	fmt.Println(isl.GetLength())
	fmt.Println(isl.BubbleSort())

	fmt.Println(isl.GetList())
}
