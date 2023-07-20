package sliceutil

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func ContainsString(array []string, element string) bool {
	for _, x := range array {
		if x == element {
			return true
		}
	}
	return false
}
