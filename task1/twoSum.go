package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, v := range nums {
		numsMap[v] = i
	}

	for i, v := range nums {
		add := target - v
		if numsMap[add] > 0 {
			return []int{i, numsMap[add]}
		}
	}

	return []int{0, 0}
}
