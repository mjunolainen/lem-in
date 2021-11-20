package routes

func LongestRoute(a []int) int {
	var temp int
	for i, val := range a {
		if i == 0 {
			temp = val
		} else {
			if val > temp {
				temp = val
			}
		}
	}
	return temp
}
