package main

// 只出现一次的数字
func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}

	for i, v := range numMap {
		if v == 1 {
			return i
		}
	}

	return -1
}
