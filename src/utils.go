package main

func isInArray(str string, arr []string) bool {
	for _, element := range arr {
		if str == element {
			return true
		}
	}
	return false
}
