package data_structs

import (
	"errors"
)

// 线性表（链式存储）
type IntLinkList struct {
	firstNode *singleNode
	length    int
}

// 初始化链表
func (ill *IntLinkList) Init() {
	ill.firstNode = &singleNode{data: 0, next: nil}
	ill.length = 0
}

// 获取链表指定位置的元素的值
func (ill *IntLinkList) GetElem(position int) (int, error) {
	if position > ill.length {
		return 0, errors.New("位置超出链表长度")
	}

	dataNode, err := ill.GetElemNode(position)

	if nil != err {
		return 0, err
	}

	return dataNode.next.data, nil
}

// 获取整个链表
func (ill *IntLinkList) GetList() (list []int) {
	dataNode := ill.firstNode.next

	for nil != dataNode {
		list = append(list, dataNode.data)
		dataNode = dataNode.next
	}

	return
}

// 获取链表长度
func (ill *IntLinkList) GetLength() int {
	return ill.length
}

// 获取链表指定位置的元素
func (ill *IntLinkList) GetElemNode(position int) (*singleNode, error) {
	if position < 1 || position-1 > ill.length {
		return nil, errors.New("索引范围溢出")
	}

	dataNode := ill.firstNode
	for ; position > 1; position-- {
		dataNode = dataNode.next
	}

	return dataNode, nil
}

// 向链表中的指定位置插入元素
func (ill *IntLinkList) InsertToIndex(position int, data []int) (err error) {
	if position < 1 || position-1 > ill.length {
		return errors.New("索引范围溢出")
	}

	if len(data) <= 0 {
		return errors.New("最少插入一个值")
	}

	prevNode, err := ill.GetElemNode(position)

	if nil != err {
		return
	}

	next := prevNode.next
	node := &singleNode{data: data[0]}
	prevNode.next = node
	ill.length++

	for _, elem := range data[1:] {
		elemNode := &singleNode{data: elem}
		node.next = elemNode
		node = elemNode
		ill.length++
	}

	node.next = next

	return
}

// 删除链表中指定位置的元素
func (ill *IntLinkList) DeleteFromIndex(position, num int) (err error) {
	if position < 1 || position-1 > ill.length {
		return errors.New("索引范围溢出")
	}

	if num <= 0 {
		return errors.New("最少删除一个值")
	}

	prevNode, err := ill.GetElemNode(position)

	if nil != err {
		return
	}

	next := prevNode.next

	for ; num > 1; num-- {
		if nil == next.next {
			break
		}

		ill.length--
		next = next.next
	}

	prevNode.next = next.next
	ill.length--

	return
}

// 链表的node结构，此为单链表
type singleNode struct {
	data int
	next *singleNode
}
