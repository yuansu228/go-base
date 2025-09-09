package main

// 有效的括号
func isValid(s string) bool {
	stack := make([]rune, 0)
	strMap := make(map[rune]rune, 3)
	strMap['('] = ')'
	strMap['['] = ']'
	strMap['{'] = '}'
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if strMap[stack[len(stack)-1]] != char {
				return false
			}
			stack = append(stack[:len(stack)-1])
		}
	}
	return len(stack) == 0
}
