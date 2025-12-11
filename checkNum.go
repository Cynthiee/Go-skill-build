package piscine

// Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

func CheckNumber(arg string) bool {
	for _, c := range arg {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}
