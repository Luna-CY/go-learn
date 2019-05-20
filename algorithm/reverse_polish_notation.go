package algorithm

import (
	"github.com/Luna-CY/go-learn/data_structs"
	"strings"
)

// 将中缀表达式转换为后缀表达式（逆波兰计算式）
func GetReversePolishNotation(value string) string {
	stack := data_structs.SortStack{}

	var result []string

main:
	for _, elem := range strings.Split(value, " ") {
		switch elem {
		case "+", "-":
			if topElem, err := stack.GetTopElem(); nil == err {
				if topElem == "*" || topElem == "/" {
					result = allStackPop(result, &stack)
				}
			}

			stack.Push(elem)
		case "*", "/", "(":
			stack.Push(elem)
		case ")":
			topElem, err := stack.Pop()
			for nil == err {
				if "(" == topElem {
					continue main
				}

				result = append(result, topElem)
				topElem, err = stack.Pop()
			}
		default:
			result = append(result, elem)
		}
	}

	result = allStackPop(result, &stack)

	return strings.Join(result, " ")
}

// 所有元素出栈
func allStackPop(buffer []string, stack *data_structs.SortStack) []string {
	stackValue, err := stack.Pop()
	for nil == err {
		buffer = append(buffer, stackValue)
		stackValue, err = stack.Pop()
	}

	return buffer
}
