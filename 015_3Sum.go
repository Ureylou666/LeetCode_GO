func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := 0
	h := true
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[k]+nums[i]+nums[j] == 0 {
					if n == 0 {
						res = append(res, []int{nums[i], nums[j], nums[k]})
						n++
					} else {
						h = true
						for m := 0; m < n; m++ {
							if res[m][0] == nums[i] && res[m][1] == nums[j] && res[m][2] == nums[k] {
								h = false
								break
							}
						}
						if h {
							res = append(res, []int{nums[i], nums[j], nums[k]})
							n++
						}
					}
				}
			}
		}
	}
	return res
}
