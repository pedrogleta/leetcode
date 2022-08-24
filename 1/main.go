package main

func twoSum(nums []int, target int) []int {

	result := make([]int, 2)

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if j, numberExists := m[target-nums[i]]; numberExists {
			if i != j {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}

	return make([]int, 0)

}
