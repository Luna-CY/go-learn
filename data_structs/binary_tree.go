package data_structs

import (
	"strconv"
	"strings"
)

type BTreeNode struct {
	data         int
	lNode, rNode *BTreeNode
}

// 获取当前节点的数据
func (tn *BTreeNode) GetData() int {
	return tn.data
}

func (tn *BTreeNode) GetLNode(new bool) *BTreeNode {
	if true == new && nil == tn.lNode {
		tn.lNode = &BTreeNode{}
	}

	return tn.lNode
}

func (tn *BTreeNode) GetRNode(new bool) *BTreeNode {
	if true == new && nil == tn.rNode {
		tn.rNode = &BTreeNode{}
	}

	return tn.rNode
}

func (tn *BTreeNode) SetData(data int) {
	tn.data = data
}

var elems []string

func CreateNewBTree(tree *BTreeNode, str string) error {
	elems = strings.Split(str, " ")

	if "#" == elems[0] {
		return nil
	}

	i, err := strconv.Atoi(elems[0])

	if nil != err {
		return err
	}

	tree.SetData(i)
	CreateNewBTree(tree.GetLNode(true), strings.Join(elems[1:], " "))
	CreateNewBTree(tree.GetRNode(true), strings.Join(elems[1:], " "))

	return nil
}

func (tn *BTreeNode) PreOrderTraverse() (res []int) {
	res = append(res, tn.data)

	if nil != tn.GetLNode(false) {
		res = append(res, tn.lNode.PreOrderTraverse()...)
	}

	if nil != tn.GetRNode(false) {
		res = append(res, tn.rNode.PreOrderTraverse()...)
	}

	return
}

func (tn *BTreeNode) InOrderTraverse() (res []int) {
	if nil != tn.GetLNode(false) {
		res = append(res, tn.lNode.InOrderTraverse()...)
	}

	res = append(res, tn.data)

	if nil != tn.GetRNode(false) {
		res = append(res, tn.rNode.InOrderTraverse()...)
	}

	return
}

func (tn *BTreeNode) PostOrderTraverse() (res []int) {
	if nil != tn.GetLNode(false) {
		res = append(res, tn.lNode.PostOrderTraverse()...)
	}

	if nil != tn.GetRNode(false) {
		res = append(res, tn.rNode.PostOrderTraverse()...)
	}

	res = append(res, tn.data)

	return
}
