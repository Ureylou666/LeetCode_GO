// 中心扩展法
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	} else {
		res := ""
		//遍历中心轴
		for i := 0; i < len(s)-1; i++ {
			// 中心轴只有一个
			res = maxPalindrome(s, i, i, res)
			res = maxPalindrome(s, i, i+1, res)
			// 中心轴有两个
		}
		return res
	}

}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(sub) > len(res) {
		return sub
	} else {
		return res
	}
}

