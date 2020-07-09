package two_sum

func twoSum(nums []int, target int) []int {
	results := make([]int, 0)
	needs := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, exist := needs[nums[i]]; exist {
			results = append(results, j, i)
		} else {
			need := target - nums[i]
			needs[need] = i
		}
	}
	return results
}
