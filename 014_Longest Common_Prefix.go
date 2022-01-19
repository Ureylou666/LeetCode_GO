func longestCommonPrefix(strs []string) string {
	prefix := ""
	temp := ""
	min := 500
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	for i := min; i > 0; i-- {
		temp = strs[0][0:i]
		h := true
		for j := 0; j < len(strs); j++ {
			if strs[j][0:i] != temp {
				h = false
				break
			}
		}
		if h {
			prefix = temp
			break
		}
	}
	return prefix
}
