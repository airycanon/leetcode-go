package container_with_most_water

func maxArea(height []int) int {
	length := len(height)
	if length == 1 {
		return 0
	}

	var (
		max   = 0
		left  = 0
		right = length - 1
	)

	for left < right {
		w := right - left
		h := 0

		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		area := w * h
		if area > max {
			max = area
		}
	}

	return max
}
