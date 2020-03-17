package number_of_steps_to_reduce_a_number_to_zero

func numberOfSteps(num int) int {
	for num > 1 {
		isEven := num%2 == 0
		if isEven {
			return 1 + numberOfSteps(int(num/2))
		} else {
			return 1 + numberOfSteps(num-1)
		}
	}

	if num == 1 {
		return 1
	}

	return 0
}
