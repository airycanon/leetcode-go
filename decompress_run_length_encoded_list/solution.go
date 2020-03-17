package decompress_run_length_encoded_list

func decompressRLElist(nums []int) []int {
	var i, length = 0, len(nums)
	var values []int

	for i < length {
		for j := nums[i]; j > 0; j-- {
			values = append(values, nums[i+1])
		}

		i += 2
	}

	return values
}
