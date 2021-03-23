package three_sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var (
		result    [][]int
		resultMap = make(map[string]struct{})
		twoSumMap = make(map[int][][]int)
		temp      = [][]int{}
	)
	for i, num := range nums {
		wanted := 0 - num
		if v, exist := twoSumMap[wanted]; exist {
			temp = v
		} else {
			temp = twoSum(nums, wanted, i)
			twoSumMap[wanted] = temp
		}

		for _, v := range temp {
			v = append(v, num)
			sort.Ints(v)
			key := fmt.Sprintf("%d%d%d", v[0], v[1], v[2])
			if _, exist := resultMap[key]; !exist {
				result = append(result, v)
				resultMap[key] = struct{}{}
			}
		}
	}

	return result
}

func twoSum(nums []int, target int, ignore int) [][]int {
	match := map[int]int{}

	var result [][]int
	for i, num := range nums {
		if i == ignore {
			continue
		}
		if v, exist := match[num]; exist {
			result = append(result, []int{num, v})
		} else {
			wanted := target - num
			match[wanted] = num
		}
	}

	return result
}
