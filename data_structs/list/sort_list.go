package list

import (
	"errors"
	"fmt"
)

// 线性表（顺序存储）
type IntSortList struct {
	list   []int
	length int
}

func (isl *IntSortList) Init(items ...int) {
	isl.length = len(items)
	isl.list = make([]int, isl.length)

	for index, item := range items {
		isl.list[index] = item
	}
}

// 检查线性表是否为空
func (isl *IntSortList) ListEmpty() bool {
	return 0 == isl.length
}

// 获取线性表
func (isl *IntSortList) GetList() []int {
	return isl.list
}

// 获取线性表长度
func (isl *IntSortList) GetLength() int {
	return isl.length
}

// 获取指定位置的元素
func (isl *IntSortList) GetElem(index int) (int, error) {
	if isl.length <= 0 {
		return -1, nil
	}

	if index < 0 || index > isl.length {
		return -1, errors.New("错误的索引位置")
	}

	return isl.list[index-1], nil
}

// 查找元素
func (isl *IntSortList) FindElem(value int) (int, error) {
	for index, elem := range isl.list {
		if elem == value {
			return index, nil
		}
	}

	return 0, errors.New("没有找到元素")
}

// 获取指定位置元素的内存地址
func (isl *IntSortList) GetElemPoint(index int) (string, error) {
	if isl.length <= 0 {
		return "", nil
	}

	if index < 0 || index > isl.length {
		return "", errors.New("错误的索引位置")
	}

	return fmt.Sprintf("%p", &isl.list[index-1]), nil
}

// 插入元素到指定索引位置
func (isl *IntSortList) InsertToIndex(position int, elem int) (err error) {
	if position < 1 || position-1 > isl.length {
		return errors.New("错误的位置")
	}

	if isl.length >= len(isl.list) {
		return errors.New("表已满")
	}

	for i := isl.length; i >= position; i-- {
		isl.list[i] = isl.list[i-1]
	}

	isl.list[position-1] = elem
	isl.length++
	return
}

// 删除指定索引位置的元素
func (isl *IntSortList) DeleteFromIndex(index int) (err error) {
	if index < 0 || index > isl.length {
		return errors.New("错误的索引位置")
	}

	if isl.length <= 0 {
		return errors.New("表内没有元素")
	}

	for ; index < isl.length; index++ {
		isl.list[index-1] = isl.list[index]
	}

	isl.list[index-1] = 0
	isl.length--
	return
}

// 交换两个元素的位置
func (isl *IntSortList) Swap(i, j int) {
	isl.list[i], isl.list[j] = isl.list[j], isl.list[i]
}

func (isl *IntSortList) BubbleSort() (count int) {
	flag := true
	for i := 0; i < isl.length && flag; i++ {
		flag = false
		for j := isl.length - 2; j >= i; j-- {
			if isl.list[j+1] < isl.list[j] {
				isl.Swap(j, j+1)
				flag = true
			}
			count++
		}
	}

	return
}
