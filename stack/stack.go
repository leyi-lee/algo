package stack

import (
	"fmt"
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == `+` || token == `-` || token == `*` || token == `/` {
			y,x := stack[len(stack)-1], stack[len(stack)-2]

			stack = stack[:len(stack)-2]
			result := calc(x, y, token)
			fmt.Println(result)
			stack = append(stack, result)
		} else {
			intToken, _ := strconv.Atoi(token)
			stack = append(stack, intToken)
		}
	}

	return stack[0]
}

func calc(x int, y int, token string) int {
	if token == "+" {
		return  x + y
	}

	if token == "-" {
		return  x - y
	}

	if token == "*" {
		return  x * y
	}

	if token == "/" {
		return  x / y
	}

	return 0
}