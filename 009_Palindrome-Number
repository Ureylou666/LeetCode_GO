func isPalindrome(x int) bool {
	var temp [256]int
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	n := 0
	for x > 0 {
		temp[n] = x % 10
		x = x / 10
		n++
	}
	n = n - 1
	i := 0
	j := n
	for temp[i] == temp[j] {
		i++
		j--
		if i > n || j < 0 {
			i--
			break
		}
	}
	if i != n {
		return false
	} else {
		return true
	}
}
