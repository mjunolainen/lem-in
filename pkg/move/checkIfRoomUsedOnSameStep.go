package move

func CheckIfRoomUsedOnSameStep(a []string, b string, c string) bool {
	if b == c {
		return false
	}
	for _, val := range a {
		if val == b {
			return true
		}
	}
	return false
}
