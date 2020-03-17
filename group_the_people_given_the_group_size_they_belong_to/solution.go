package group_the_people_given_the_group_size_they_belong_to

func groupThePeople(groupSizes []int) [][]int {
	result := make([][]int, 0, 0)
	groupSizeMap := map[int][]int{}

	for id, groupSize := range groupSizes {
		_, exist := groupSizeMap[groupSize]

		if !exist {
			groupSizeMap[groupSize] = []int{id}

		} else if len(groupSizeMap[groupSize]) < groupSize {
			groupSizeMap[groupSize] = append(groupSizeMap[groupSize], id)
		}

		if len(groupSizeMap[groupSize]) == groupSize {
			result = append(result, groupSizeMap[groupSize])
			groupSizeMap[groupSize] = []int{}
		}
	}

	return result
}
