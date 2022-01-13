func myAtoi(s string) int {
	output := 0
	fushu := false
	fuhao := 0
	i := 0
	for i < len(s) {
		for s[i] == ' ' {
			i++
			if i == len(s) {
				return 0
			}
		}
		if s[i] != ' ' && s[i] != '+' && s[i] != '-' && (s[i] < '0' || s[i] > '9') {
			return 0
		}
		for s[i] == '+' || s[i] == '-' {
			fuhao++
			if s[i] == '-' {
				fushu = true
			}
			i++
			if i == len(s) {
				return 0
			}
		}
		if fuhao > 1 {
			return 0
		}
		for s[i] >= '0' && s[i] <= '9' {
			output = output * 10
			output = output + int(s[i]) - int('0')
			if output > 1<<31-1 || output < -(1<<31) {
				if fushu {
					return -2147483648
				} else {
					return 2147483647
				}
			}
			i++
			if i == len(s) {
				break
			}
		}
		if fushu {
			output = -output
		}
		return output
	}
	return output
}
