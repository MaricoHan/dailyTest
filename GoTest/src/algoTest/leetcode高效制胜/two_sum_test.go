package leetcode

// 时间复杂度 O(n)
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int) // 存储历史的所有值： num->index
	for i, num := range nums {
		if index, ok := tmp[target-num]; ok { // 看下是否已存有契合的数字
			return []int{index, i}
		}
		tmp[num] = i
	}

	return []int{}
}
