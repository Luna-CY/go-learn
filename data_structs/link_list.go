package data_structs

import "errors"

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
func (ill *IntLinkList) GetElem(index int) (int, error) {
	dataNode, err := ill.GetElemNode(index)

	if err != nil {
		return 0, err
	}

	return dataNode.data, nil
}

// 获取链表指定位置的元素
func (ill *IntLinkList) GetElemNode(index int) (*singleNode, error) {
	if index < 1 || index-1 > ill.length {
		return nil, errors.New("索引范围溢出")
	}

	dataNode := ill.firstNode.next
	for ; index > 1; index-- {
		dataNode = dataNode.next
	}

	return dataNode, nil
}

// 向链表中的指定位置插入元素
func (ill *IntLinkList) InsertToIndex(index int, data []int) (err error) {
	if index < 1 || index-1 > ill.length {
		return errors.New("索引范围溢出")
	}

	if len(data) <= 0 {
		return errors.New("最少插入一个值")
	}

	prevNode, err := ill.GetElemNode(index)

	if err != nil {
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

// 链表的node结构，此为单链表
type singleNode struct {
	data int
	next *singleNode
}
