package data_structs

import "errors"

type SortStack struct {
	stack  []string
	length int
}

// 推入一个元素
func (ss *SortStack) Push(elem string) {
	ss.stack = append(ss.stack, elem)
	ss.length++
}

// 弹出头元素
func (ss *SortStack) Pop() (elem string, err error) {
	if 0 >= ss.length {
		return "", errors.New("空栈无法获取元素")
	}

	elem = ss.stack[ss.length-1]
	ss.stack = ss.stack[:ss.length-1]
	ss.length--

	return
}

func (ss *SortStack) GetTopElem() (string, error) {
	if 0 >= ss.length {
		return "", errors.New("空栈无法获取元素")
	}

	return ss.stack[ss.length-1], nil
}
