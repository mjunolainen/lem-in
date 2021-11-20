package routes

func CompareTwoRoutes(a []string, b []string) bool {
	for _, val := range a {
		for _, val2 := range b {
			if val == val2 {
				return false
			}
		}
	}
	return true
}
