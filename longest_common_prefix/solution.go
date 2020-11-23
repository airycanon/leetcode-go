package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	find := ""
	if len(strs) == 0 {
		return find
	}

	for i := 0; i < len(strs[0]); i++ {
		target := []rune(strs[0])[i]
		for _, str := range strs {
			if len(str) < i+1 {
				return find
			}

			if []rune(str)[i] != target {
				return find
			}
		}
		find += string(target)
	}

	return find
}
