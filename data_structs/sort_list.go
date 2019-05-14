package data_structs

import (
	"errors"
	"fmt"
)

// 线性表（顺序存储）
type IntSortList struct {
	list   [15]int
	length int
}

func (sl *IntSortList) GetList() [15]int {
	return sl.list
}

// 获取线性表长度
func (sl *IntSortList) GetLength() int {
	return sl.length
}

// 获取指定位置的元素
func (sl *IntSortList) GetElem(index int) (int, error) {
	if sl.length <= 0 {
		return -1, nil
	}

	if index < 0 || index > sl.length {
		return -1, errors.New("错误的索引位置")
	}

	return sl.list[index-1], nil
}

// 获取指定位置元素的内存地址
func (sl *IntSortList) GetElemPoint(index int) (string, error) {
	if sl.length <= 0 {
		return "", nil
	}

	if index < 0 || index > sl.length {
		return "", errors.New("错误的索引位置")
	}

	return fmt.Sprintf("%p", &sl.list[index-1]), nil
}

// 插入元素到指定索引位置
func (sl *IntSortList) InsertToIndex(index int, elem int) (err error) {
	if index < 1 || index-1 > sl.length {
		return errors.New("错误的索引位置")
	}

	if sl.length >= len(sl.list) {
		return errors.New("表已满")
	}

	for ; index <= sl.length; index++ {
		ce := sl.list[index-1]
		sl.list[index-1] = elem
		elem = ce
	}

	sl.list[index-1] = elem
	sl.length += 1
	return
}

// 删除指定索引位置的元素
func (sl *IntSortList) DeleteFromIndex(index int) (err error) {
	if index < 0 || index > sl.length {
		return errors.New("错误的索引位置")
	}

	if sl.length <= 0 {
		return errors.New("表内没有元素")
	}

	for ; index < sl.length; index++ {
		sl.list[index-1] = sl.list[index]
	}

	sl.list[index-1] = 0
	sl.length--
	return
}
