package main

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := []rune(strs[0])
	for _, word := range strs {
		for i, char := range word {
			if len(prefix) > 0 && char != prefix[i] {
				prefix = append(prefix[:i])
				break
			}
		}
	}
	return string(prefix)
}
