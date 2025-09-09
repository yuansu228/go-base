package main

// 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x > 0 && x%10 == 0 {
		return false
	}

	revert := 0
	origin := x

	for x > 0 {
		revert = revert*10 + x%10
		x /= 10
	}

	return origin == revert

}
