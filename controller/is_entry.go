package controller

func is_entry(arr *[]int, num int) bool {
	for _, v := range *arr {
		if v == num {
			return true
		}
	}

	return false
}
