package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	arr := make([]string, 0)
	dfs(n, n, "", n, &arr)
	return arr
}

func dfs(leftBranch int, rightBranch int, str string, n int, res *[]string) {
	fmt.Println("leftBranch=>", leftBranch, "  rightBranch=>", rightBranch, str, n)
	if len(str) == 2*n {
		// 如果长度匹配则不用继续寻找直接放进数组中\
		*res = append(*res, str)
		return
	}

	// 左括号有剩下，才可以选他
	if leftBranch > 0 {
		dfs(leftBranch-1, rightBranch, str+"(", n, res)
	}

	// 右括号比左括号多，才能选择右括号
	if leftBranch < rightBranch {
		dfs(leftBranch, rightBranch-1, str+")", n, res)
	}
}
