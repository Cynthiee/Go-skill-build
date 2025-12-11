package piscine

// Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.

func CountAlpha(s string) int {
	count := 0
	for _, a := range s {
		if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') {
			count++
		}
	}
	return count
}
