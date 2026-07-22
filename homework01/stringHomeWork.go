package main

import "fmt"

func main() {
	fmt.Println(isValid("{}[]"))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	// 如果字符串的长度是奇数，肯定不符合匹配规则
	if len(s)%2 != 0 {
		return false
	}
	// 用切片来模拟栈结构，存放btye类型的字符
	stack := []byte{}
	// stack := make([]byte, 0 ,len(s))
	// var stack []byte 声明一个nil切片
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' {
			stack = append(stack, ')')
		} else if char == '{' {
			stack = append(stack, '}')
		} else if char == '[' {
			stack = append(stack, ']')
		} else {
			// 此时遇见了右括号，如果栈为空，说明这个右括号是多余的，或者栈顶的元素不匹配，直接返回false
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			// 如果栈顶的元素匹配，则弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0 // 如果栈为空，说明所有括号都匹配，返回true
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串的长度小于i，或者当前字符串的第i个字符不等于indexZeroStr的第i个字符，则返回indexZeroStr的前i个字符
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
