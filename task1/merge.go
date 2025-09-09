package main

// 合并区间
func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][0] > intervals[i+1][0] {
			temp := intervals[i]
			intervals[i] = intervals[i+1]
			intervals[i+1] = temp
		}
	}

	result := [][]int{intervals[0]}
	last := len(result) - 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > result[last][1] {
			result = append(result, intervals[i])
		} else {
			if intervals[i][1] > result[last][1] {
				result[last][1] = intervals[i][1]
			}
		}
	}

	return result
}
