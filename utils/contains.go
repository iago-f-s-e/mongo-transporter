package utils

func ContainsInArrString(arr []string, compare string) bool {
	for _, x := range arr {
		if x == compare {
			return true
		}
	}

	return false
}
