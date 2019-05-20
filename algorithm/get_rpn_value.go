package algorithm

import (
	"errors"
	"github.com/Luna-CY/go-learn/data_structs"
	"strconv"
	"strings"
)

func GetRPNValue(rpn string) (float32, error) {
	stack := data_structs.SortStack{}

	for _, elem := range strings.Split(rpn, " ") {
		switch elem {
		case "+":
			i1, i2, err := getOperateElems(&stack)

			if nil != err {
				return 0.0, err
			}

			stack.Push(itoa(i2 + i1))
		case "-":
			i1, i2, err := getOperateElems(&stack)

			if nil != err {
				return 0.0, err
			}

			stack.Push(itoa(i2 - i1))
		case "*":
			i1, i2, err := getOperateElems(&stack)

			if nil != err {
				return 0.0, err
			}

			stack.Push(itoa(i2 * i1))
		case "/":
			i1, i2, err := getOperateElems(&stack)

			if nil != err {
				return 0.0, err
			}

			stack.Push(itoa(i2 / i1))
		default:
			stack.Push(elem)
		}
	}

	result, err := stack.Pop()

	if nil != err {
		return 0.0, err
	}

	// 这里会丢失小数位
	return float32(atoi(result)), nil
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)

	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func getOperateElems(stack *data_structs.SortStack) (int, int, error) {
	i1, err := stack.Pop()

	if nil != err {
		return 0, 0, errors.New("计算式格式错误")
	}

	i2, err := stack.Pop()

	if nil != err {
		return 0, 0, errors.New("计算式格式错误")
	}

	return atoi(i1), atoi(i2), nil
}
