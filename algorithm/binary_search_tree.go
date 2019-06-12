package algorithm

type BSTree struct {
	data           int
	lChild, rChild *BSTree
}

func (bst *BSTree) Search(key int) (*BSTree, bool) {

	if key == bst.data {
		return bst, true
	}

	if key < bst.data && nil != bst.lChild {
		return bst.lChild.Search(key)
	}

	if key > bst.data && nil != bst.rChild {
		return bst.rChild.Search(key)
	}

	return bst, false
}

func (bst *BSTree) Insert(key int) bool {
	if parent, ok := bst.Search(key); false == ok {
		if 0 == parent.data {
			parent.data = key

			return true
		}

		node := &BSTree{data: key}

		if key < (*parent).data {
			parent.lChild = node
		} else {
			parent.rChild = node
		}

		return true
	}

	return false
}

func (bst *BSTree) GetLNode(new bool) *BSTree {
	if true == new && nil == bst.lChild {
		bst.lChild = &BSTree{}
	}

	return bst.lChild
}

func (bst *BSTree) GetRNode(new bool) *BSTree {
	if true == new && nil == bst.rChild {
		bst.rChild = &BSTree{}
	}

	return bst.rChild
}

func (bst *BSTree) PreOrderTraverse() (res []int) {
	res = append(res, bst.data)

	if nil != bst.GetLNode(false) {
		res = append(res, bst.lChild.PreOrderTraverse()...)
	}

	if nil != bst.GetRNode(false) {
		res = append(res, bst.rChild.PreOrderTraverse()...)
	}

	return
}

func (bst *BSTree) InOrderTraverse() (res []int) {
	if nil != bst.GetLNode(false) {
		res = append(res, bst.lChild.InOrderTraverse()...)
	}

	res = append(res, bst.data)

	if nil != bst.GetRNode(false) {
		res = append(res, bst.rChild.InOrderTraverse()...)
	}

	return
}
