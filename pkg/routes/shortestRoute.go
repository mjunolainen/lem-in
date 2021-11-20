package routes

func ShortestRoute(a []int) int {
	var temp int
	var temp2 int
	for i, val := range a {
		if i == 0 {
			temp = val
			temp2 = i
		} else if val < temp {
			temp = val
			temp2 = i
		}
	}
	return temp2
}
