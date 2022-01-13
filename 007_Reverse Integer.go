func reverse(x int) int {
	output := 0
	for x != 0 {
		output = output + x%10
		x = x / 10
		if x != 0 {
			output = output * 10
		}
	}
	if output > 1<<31-1 || output < -(1<<31) {
		output = 0
	}
	return output
}

func main() {
	output := reverse(1534236469)
	println(output)
}
